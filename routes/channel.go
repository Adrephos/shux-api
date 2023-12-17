package routes

import (
	"time"
	"github.com/gofiber/fiber/v2"
	"github.com/shuxbot/shux-api/infrastructure/routing"
	"github.com/shuxbot/shux-api/middlewares"
)

func ChannelEndpoints(app *fiber.App, routeHandler *routing.RouteHandler, ttl time.Duration) {
	//Channel endpoints
	route := app.Group("api/v1/servers/:server_id")

	route.Get("/channels", middlewares.VerifyToken, middlewares.CacheAdd(ttl), routeHandler.ListChannels)
	route.Get("/channels/:channel_id", middlewares.VerifyToken, middlewares.CacheAdd(ttl), routeHandler.GetChannel)
	route.Delete("/channels/:channel_id", middlewares.VerifyToken, middlewares.CacheAdd(ttl), routeHandler.DeleteChannel)
	route.Patch("/channels/:channel_id", middlewares.VerifyToken, middlewares.CacheAdd(ttl), routeHandler.UpdateChannel)
	route.Put("/channels/:channel_id", middlewares.VerifyToken, middlewares.CacheAdd(ttl), routeHandler.ReplaceChannel)
	route.Post("/channels/:channel_id", middlewares.VerifyToken, middlewares.CacheAdd(ttl), routeHandler.CreateChannel)
}
