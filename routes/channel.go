package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shuxbot/shux-api/infrastructure/routing"
	"github.com/shuxbot/shux-api/infrastructure/transient"
)

func ChannelEndpoints(app *fiber.App, routeHandler *routing.RouteHandler) {
	//Channel endpoints
	route := app.Group("api/v1/servers/:server_id")

	route.Get("/channels", transient.New(false), routeHandler.ListChannels)
	route.Get("/channels/:channel_id", transient.New(true), routeHandler.GetChannel)
	route.Delete("/channels/:channel_id", transient.New(true), routeHandler.DeleteChannel)
	route.Patch("/channels/:channel_id", transient.New(true), routeHandler.UpdateChannel)
	route.Put("/channels/:channel_id", transient.New(true), routeHandler.ReplaceChannel)
	route.Post("/channels/:channel_id", transient.New(true), routeHandler.CreateChannel)
}
