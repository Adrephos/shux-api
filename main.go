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
	userApp := application.NewUserRepo(userRepo)
	routeHandler := infrastructure.NewRouteHandler(userApp)

	app := fiber.New()

	app.Get("/servers/:server_id/users/:user_id", routeHandler.GetUser)
	app.Delete("/servers/:server_id/users/:user_id", routeHandler.DeleteUser)
	app.Post("/servers/:server_id/users/:user_id", routeHandler.CreateUser)

	app.Listen(":3000")
}
