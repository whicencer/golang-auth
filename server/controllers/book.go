package controllers

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/whicencer/golang-auth/database"
	"github.com/whicencer/golang-auth/models"
	"gorm.io/gorm"
)

// NewBook Create a new book
func NewBook(c *fiber.Ctx) error {
	db := database.DB.Db
	book := new(models.Book)

	err := c.BodyParser(book)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "Something's wrong with your input",
			"data":    nil,
		})
	}

	err = db.Create(&book).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "Could not create a book",
			"data":    err,
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"status":  "success",
		"message": "Book has created",
		"data":    book,
	})
}

// GetBooks Get all books
func GetBooks(c *fiber.Ctx) error {
	db := database.DB.Db
	var books []models.Book

	db.Find(&books)

	return c.Status(200).JSON(fiber.Map{
		"status":  "sucess",
		"message": "Books Found",
		"data":    books,
	})
}

// GetBook Get Single Book
func GetBook(c *fiber.Ctx) error {
	db := database.DB.Db
	id := c.Params("id")
	var book models.Book

	result := db.Where("id = ?", id).First(&book)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Book not found",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "Book found",
		"data":    book,
	})
}

// DeleteBook Delete Book
func DeleteBook(c *fiber.Ctx) error {
	db := database.DB.Db
	id := c.Params("id")
	var book models.Book

	result := db.Delete(&book, id)
	if result.RowsAffected == 0 {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "Book was not found",
			"data":    nil,
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "Book was successfully deleted",
		"data":    book,
	})
}
