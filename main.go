package main

import (
	"os"
	"time"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"github.com/shuxbot/shux-api/application"
	"github.com/shuxbot/shux-api/infrastructure"
	"github.com/shuxbot/shux-api/infrastructure/persistance"
	"github.com/shuxbot/shux-api/infrastructure/routing"
	"github.com/shuxbot/shux-api/routes"
)

func main() {
	// Firestore client
	firestoreClient := persistance.Client

	// Initialize all repos and apps
	userRepo := infrastructure.NewFirestoreUserRepo(firestoreClient)
	channelRepo := infrastructure.NewFirestoreChannelRepo(firestoreClient)
	roleRepo := infrastructure.NewFirestoreRoleRepo(firestoreClient)
	serverRepo := infrastructure.NewFirestoreServerRepo(firestoreClient)
	adminRepo := infrastructure.NewFirestoreAdminRepo(firestoreClient)

	userApp := application.NewUserApp(userRepo)
	channelApp := application.NewChannelApp(channelRepo)
	roleApp := application.NewRoleApp(roleRepo)
	serverApp := application.NewServerApp(serverRepo)
	adminApp := application.NewAdminApp(adminRepo)

	// Initialize route handler
	routeHandler := routing.NewRouteHandler(
		userApp,
		channelApp,
		roleApp,
		serverApp,
		adminApp,
	)

	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	app.Use(logger.New())

	ttl := 30 * time.Minute
	routes.UserEndpoints(app, routeHandler, ttl)
	routes.ChannelEndpoints(app, routeHandler, ttl)
	routes.RoleEndpoints(app, routeHandler, ttl)
	routes.ServerEndpoints(app, routeHandler, ttl)
	routes.JWTEndpoints(app, routeHandler)

	app.Listen(":" + os.Getenv("API_PORT"))
}
