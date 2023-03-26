package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/whicencer/golang-auth/database"
	"github.com/whicencer/golang-auth/routes"
)

func main() {
	app := fiber.New()

	// Connecting Database
	database.Connect()

	routes.SetupRoutes(app)

	app.Listen("localhost:2000")
}
