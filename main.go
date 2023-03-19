package main

import (
	"github.com/whicencer/react-finance-server/routes"

	"github.com/gofiber/fiber/v2"
)

type DataStruct struct {
	Name string
	Age  int
}
type ResponseStruct struct {
	Data DataStruct
}

func main() {
	app := fiber.New()

	routes.SetupRoutes(app)

	app.Listen("localhost:3000")
}
