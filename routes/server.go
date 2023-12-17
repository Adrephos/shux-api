package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shuxbot/shux-api/infrastructure/routing"
	"github.com/shuxbot/shux-api/middlewares"
	"time"
)

func ServerEndpoints(app *fiber.App, routeHandler *routing.RouteHandler, ttl time.Duration) {
	//Server endpoints
	route := app.Group("api/v1")

	route.Get("/servers", middlewares.VerifyToken, routeHandler.ListServers)
	route.Get("/servers/:server_id/leaderboard", middlewares.VerifyToken, middlewares.CacheAdd(ttl), routeHandler.ServerRanking)
}
