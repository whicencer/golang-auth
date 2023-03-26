package user

import (
	"github.com/gofiber/fiber/v2"
	"github.com/whicencer/golang-auth/controllers"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/users", controllers.GetUsers)          // Get all users
	app.Get("/users/:nickname", controllers.GetUser) // Get a single user
}
