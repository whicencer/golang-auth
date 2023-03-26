package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/whicencer/golang-auth/routes/auth"
	"github.com/whicencer/golang-auth/routes/book"
	"github.com/whicencer/golang-auth/routes/user"
)

func SetupRoutes(app *fiber.App) {
	book.SetupRoutes(app)
	user.SetupRoutes(app)
	auth.SetupRoutes(app)
}
