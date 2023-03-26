package controllers

import (
	"github.com/gofiber/fiber/v2"
)

// GetUsers Get all users
func GetUsers(c *fiber.Ctx) error {
	return c.SendString("Get all users")
}

// GetUser Get a single user
func GetUser(c *fiber.Ctx) error {
	return c.SendString("Get a single user")
}
