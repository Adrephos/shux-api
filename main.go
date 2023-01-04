package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shuxbot/shux-api/application"
	"github.com/shuxbot/shux-api/infrastructure"
	"github.com/shuxbot/shux-api/infrastructure/persistance"
)

func main() {

	firestoreClient := persistance.Client
	userRepo := infrastructure.NewFirestoreUserRepo(firestoreClient)
	channelRepo := infrastructure.NewFirestoreChannelRepo(firestoreClient)
	userApp := application.NewUserApp(userRepo)
	channelApp := application.NewChannelApp(channelRepo)
	routeHandler := infrastructure.NewRouteHandler(userApp, channelApp)

	app := fiber.New()

	//User endpoints
	app.Get("/servers/:server_id/users/:user_id", routeHandler.GetUser)
	app.Delete("/servers/:server_id/users/:user_id", routeHandler.DeleteUser)
	app.Patch("/servers/:server_id/users/:user_id", routeHandler.UpdateUser)
	app.Put("/servers/:server_id/users/:user_id", routeHandler.CreateUser)
	app.Post("/servers/:server_id/users/:user_id", routeHandler.CreateUser)

	//Channel endpoints
	app.Get("/servers/:server_id/channels", routeHandler.ListChannels)

	app.Listen(":3000")
}
