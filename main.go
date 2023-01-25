package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/shuxbot/shux-api/application"
	"github.com/shuxbot/shux-api/infrastructure"
	"github.com/shuxbot/shux-api/infrastructure/persistance"
	"github.com/shuxbot/shux-api/infrastructure/transient"
	"github.com/goccy/go-json"
)

func main() {

	//Initialize all repos and apps
	firestoreClient := persistance.Client
	userRepo := infrastructure.NewFirestoreUserRepo(firestoreClient)
	channelRepo := infrastructure.NewFirestoreChannelRepo(firestoreClient)
	roleRepo := infrastructure.NewFirestoreRoleRepo(firestoreClient)
	serverRepo := infrastructure.NewFirestoreServerRepo(firestoreClient)
	userApp := application.NewUserApp(userRepo)
	channelApp := application.NewChannelApp(channelRepo)
	roleApp := application.NewRoleApp(roleRepo)
	serverApp := application.NewServerApp(serverRepo)
	routeHandler := infrastructure.NewRouteHandler(userApp, channelApp, roleApp, serverApp)

	app := fiber.New(fiber.Config{
        JSONEncoder: json.Marshal,
        JSONDecoder: json.Unmarshal,
    })

	app.Use(logger.New())

	//Server endpoints
	app.Get("/servers", transient.New(false), routeHandler.ListServers)
	app.Get("/servers/:server_id/ranking", transient.New(false), routeHandler.ServerRanking)
	app.Get("/servers/:server_id/ranking/:user_id", transient.New(false), routeHandler.ServerUserRanking)

	//User endpoints
	app.Get("/servers/:server_id/users/:user_id", transient.New(true), routeHandler.GetUser)
	app.Delete("/servers/:server_id/users/:user_id", transient.New(true), routeHandler.DeleteUser)
	app.Patch("/servers/:server_id/users/:user_id", transient.New(true), routeHandler.UpdateUser)
	app.Put("/servers/:server_id/users/:user_id", transient.New(true), routeHandler.ReplaceUser)
	app.Post("/servers/:server_id/users/:user_id", transient.New(true), routeHandler.CreateUser)

	//Channel endpoints
	app.Get("/servers/:server_id/channels", transient.New(false), routeHandler.ListChannels)
	app.Get("/servers/:server_id/channels/:channel_id", transient.New(true), routeHandler.GetChannel)
	app.Delete("/servers/:server_id/channels/:channel_id", transient.New(true), routeHandler.DeleteChannel)
	app.Patch("/servers/:server_id/channels/:channel_id", transient.New(true), routeHandler.UpdateChannel)
	app.Put("/servers/:server_id/channels/:channel_id", transient.New(true), routeHandler.ReplaceChannel)
	app.Post("/servers/:server_id/channels/:channel_id", transient.New(true), routeHandler.CreateChannel)

	//Role endpoints
	app.Get("/servers/:server_id/roles", transient.New(false), routeHandler.ListRoles)
	app.Get("/servers/:server_id/roles/:role_id", transient.New(true), routeHandler.GetRole)
	app.Delete("/servers/:server_id/roles/:role_id", transient.New(true), routeHandler.DeleteRole)
	app.Patch("/servers/:server_id/roles/:role_id", transient.New(true), routeHandler.UpdateRole)
	app.Put("/servers/:server_id/roles/:role_id", transient.New(true), routeHandler.ReplaceRole)
	app.Post("/servers/:server_id/roles/:role_id", transient.New(true), routeHandler.CreateRole)

	app.Listen(":3000")
}
