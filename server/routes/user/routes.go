package user

import (
	"github.com/gofiber/fiber/v2"
	"github.com/whicencer/golang-auth/controllers"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/users/:nickname", controllers.GetUser)
	app.Get("/users", func(c *fiber.Ctx) error {
		return c.Redirect("/users/:nickname")
	})
}
