package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shuxbot/shux-api/infrastructure/routing"
	"github.com/shuxbot/shux-api/middlewares"
)

func ChannelEndpoints(app *fiber.App, routeHandler *routing.RouteHandler) {
	//Channel endpoints
	route := app.Group("v1/servers/:server_id")

	route.Get("/channels", middlewares.VerifyToken, middlewares.NewInCache(false), routeHandler.ListChannels)
	route.Get("/channels/:channel_id", middlewares.VerifyToken, middlewares.NewInCache(true), routeHandler.GetChannel)
	route.Delete("/channels/:channel_id", middlewares.VerifyToken, middlewares.NewInCache(true), routeHandler.DeleteChannel)
	route.Patch("/channels/:channel_id", middlewares.VerifyToken, middlewares.NewInCache(true), routeHandler.UpdateChannel)
	route.Put("/channels/:channel_id", middlewares.VerifyToken, middlewares.NewInCache(true), routeHandler.ReplaceChannel)
	route.Post("/channels/:channel_id", middlewares.VerifyToken, middlewares.NewInCache(true), routeHandler.CreateChannel)
}
