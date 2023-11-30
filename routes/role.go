package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shuxbot/shux-api/infrastructure/routing"
	"github.com/shuxbot/shux-api/middlewares"
)

func RoleEndpoints(app *fiber.App, routeHandler *routing.RouteHandler) {
	//Role endpoints
	route := app.Group("v1/servers/:server_id")

	route.Get("/roles", middlewares.VerifyToken, middlewares.NewInCache(false), routeHandler.ListRoles)
	route.Get("/roles/:role_id", middlewares.VerifyToken, middlewares.NewInCache(true), routeHandler.GetRole)
	route.Delete("/roles/:role_id", middlewares.VerifyToken, middlewares.NewInCache(true), routeHandler.DeleteRole)
	route.Patch("/roles/:role_id", middlewares.VerifyToken, middlewares.NewInCache(true), routeHandler.UpdateRole)
	route.Put("/roles/:role_id", middlewares.VerifyToken, middlewares.NewInCache(true), routeHandler.ReplaceRole)
	route.Post("/roles/:role_id", middlewares.VerifyToken, middlewares.NewInCache(true), routeHandler.CreateRole)
}
