package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shuxbot/shux-api/infrastructure/routing"
	"github.com/shuxbot/shux-api/middlewares"
)

func JWTEndpoints(app *fiber.App, routeHandler *routing.RouteHandler) {
	route := app.Group("api/v1/auth")

	route.Post("refresh", routeHandler.RefreshToken)
	route.Post("login", routeHandler.Login)
	route.Post("register", middlewares.RegisterKey, routeHandler.Register)
}
