package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shuxbot/shux-api/infrastructure/routing"
	"github.com/shuxbot/shux-api/infrastructure/transient"
)

func RoleEndpoints(app *fiber.App, routeHandler *routing.RouteHandler) {
	//Role endpoints
	route := app.Group("api/v1/servers/:server_id")

	route.Get("/roles", transient.New(false), routeHandler.ListRoles)
	route.Get("/roles/:role_id", transient.New(true), routeHandler.GetRole)
	route.Delete("/roles/:role_id", transient.New(true), routeHandler.DeleteRole)
	route.Patch("/roles/:role_id", transient.New(true), routeHandler.UpdateRole)
	route.Put("/roles/:role_id", transient.New(true), routeHandler.ReplaceRole)
	route.Post("/roles/:role_id", transient.New(true), routeHandler.CreateRole)
}
