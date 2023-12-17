package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shuxbot/shux-api/infrastructure/routing"
	"github.com/shuxbot/shux-api/middlewares"
	"time"
)

func RoleEndpoints(app *fiber.App, routeHandler *routing.RouteHandler, ttl time.Duration) {
	//Role endpoints
	route := app.Group("api/v1/servers/:server_id")

	route.Get("/roles", middlewares.VerifyToken, middlewares.CacheAdd(ttl), routeHandler.ListRoles)
	route.Get("/roles/:role_id", middlewares.VerifyToken, middlewares.CacheAdd(ttl), routeHandler.GetRole)
	route.Delete("/roles/:role_id", middlewares.VerifyToken, middlewares.CacheAdd(ttl), routeHandler.DeleteRole)
	route.Patch("/roles/:role_id", middlewares.VerifyToken, middlewares.CacheAdd(ttl), routeHandler.UpdateRole)
	route.Put("/roles/:role_id", middlewares.VerifyToken, middlewares.CacheAdd(ttl), routeHandler.ReplaceRole)
	route.Post("/roles/:role_id", middlewares.VerifyToken, middlewares.CacheAdd(ttl), routeHandler.CreateRole)
}
