package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shuxbot/shux-api/infrastructure/routing"
	"github.com/shuxbot/shux-api/infrastructure/transient"
)

func UserEndpoints(app *fiber.App, routeHandler *routing.RouteHandler) {
	//User endpoints
	route := app.Group("api/v1/servers/:server_id")

	route.Get("/users/:user_id", transient.New(true), routeHandler.GetUser)
	route.Delete("/users/:user_id", transient.New(true), routeHandler.DeleteUser)
	route.Patch("/users/:user_id", transient.New(true), routeHandler.UpdateUser)
	route.Put("/users/:user_id", transient.New(true), routeHandler.ReplaceUser)
	route.Post("/users/:user_id", transient.New(true), routeHandler.CreateUser)
}
