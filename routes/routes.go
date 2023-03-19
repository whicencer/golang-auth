package routes

import (
	"github.com/whicencer/react-finance-server/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/api", controllers.Hello)
	app.Get("/api/test", controllers.Test)
}
