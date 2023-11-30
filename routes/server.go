package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shuxbot/shux-api/infrastructure/routing"
	"github.com/shuxbot/shux-api/middlewares"
)

func ServerEndpoints(app *fiber.App, routeHandler *routing.RouteHandler) {
	//Server endpoints
	route := app.Group("v1")

	route.Get("/servers", middlewares.VerifyToken, routeHandler.ListServers)
	route.Get("/servers/:server_id/leaderboard", middlewares.VerifyToken, middlewares.NewInCache(false), routeHandler.ServerRanking)
}
