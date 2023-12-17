package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shuxbot/shux-api/infrastructure/routing"
	"github.com/shuxbot/shux-api/middlewares"
	"time"
)

func UserEndpoints(app *fiber.App, routeHandler *routing.RouteHandler, ttl time.Duration) {
	//User endpoints
	route := app.Group("api/v1/servers/:server_id")

	route.Get("/users/:user_id", middlewares.VerifyToken, middlewares.CacheAdd(ttl), routeHandler.GetUser)
	route.Delete("/users/:user_id", middlewares.VerifyToken, middlewares.CacheAdd(ttl), routeHandler.DeleteUser)
	route.Patch("/users/:user_id", middlewares.VerifyToken, middlewares.CacheAdd(ttl), routeHandler.UpdateUser)
	route.Put("/users/:user_id", middlewares.VerifyToken, middlewares.CacheAdd(ttl), routeHandler.ReplaceUser)
	route.Post("/users/:user_id", middlewares.VerifyToken, middlewares.CacheAdd(ttl), routeHandler.CreateUser)
	route.Get("/users/:user_id/rank", middlewares.VerifyToken, middlewares.CacheAdd(ttl), routeHandler.ServerUserRanking)
}
