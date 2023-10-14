package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shuxbot/shux-api/infrastructure/routing"
	"github.com/shuxbot/shux-api/infrastructure/transient"
)

func ServerEndpoints(app *fiber.App, routeHandler *routing.RouteHandler) {
	//Server endpoints
	route := app.Group("api/v1")

	route.Get("/servers", transient.New(false), routeHandler.ListServers)
	route.Get("/servers/:server_id/ranking", transient.New(false), routeHandler.ServerRanking)
	route.Get("/servers/:server_id/ranking/:user_id", transient.New(false), routeHandler.ServerUserRanking)
}
