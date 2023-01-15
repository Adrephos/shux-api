package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shuxbot/shux-api/application"
	"github.com/shuxbot/shux-api/infrastructure"
	"github.com/shuxbot/shux-api/infrastructure/persistance"
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

	app := fiber.New()

	//Server endpoints
	app.Get("/servers", routeHandler.ListServers)
	app.Get("/servers/:server_id/ranking", routeHandler.ServerRanking)

	//User endpoints
	app.Get("/servers/:server_id/users/:user_id", routeHandler.GetUser)
	app.Delete("/servers/:server_id/users/:user_id", routeHandler.DeleteUser)
	app.Patch("/servers/:server_id/users/:user_id", routeHandler.UpdateUser)
	app.Put("/servers/:server_id/users/:user_id", routeHandler.ReplaceUser)
	app.Post("/servers/:server_id/users/:user_id", routeHandler.CreateUser)

	//Channel endpoints
	app.Get("/servers/:server_id/channels", routeHandler.ListChannels)
	app.Get("/servers/:server_id/channels/:channel_id", routeHandler.GetChannel)
	app.Delete("/servers/:server_id/channels/:channel_id", routeHandler.DeleteChannel)
	app.Patch("/servers/:server_id/channels/:channel_id", routeHandler.UpdateChannel)
	app.Put("/servers/:server_id/channels/:channel_id", routeHandler.ReplaceChannel)
	app.Post("/servers/:server_id/channels/:channel_id", routeHandler.CreateChannel)

	//Role endpoints
	app.Get("/servers/:server_id/roles", routeHandler.ListRoles)
	app.Get("/servers/:server_id/roles/:role_id", routeHandler.GetRole)
	app.Delete("/servers/:server_id/roles/:role_id", routeHandler.DeleteRole)
	app.Patch("/servers/:server_id/roles/:role_id", routeHandler.UpdateRole)
	app.Put("/servers/:server_id/roles/:role_id", routeHandler.ReplaceRole)
	app.Post("/servers/:server_id/roles/:role_id", routeHandler.CreateRole)

	app.Listen(":3000")
}
