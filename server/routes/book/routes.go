package book

import (
	"github.com/gofiber/fiber/v2"
	"github.com/whicencer/golang-auth/controllers"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/books", controllers.GetBooks)
	app.Get("/books/:id", controllers.GetBook)
	app.Post("/books", controllers.NewBook)
	app.Delete("/books/:id", controllers.DeleteBook)
}
