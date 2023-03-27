package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/whicencer/golang-auth/database"
	"github.com/whicencer/golang-auth/routes"
)

func main() {
	app := fiber.New()

	app.Use(cors.New())

	// Connecting Database
	database.Connect()

	routes.SetupRoutes(app)

	app.Listen("localhost:2000")
}
