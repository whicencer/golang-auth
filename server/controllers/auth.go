package controllers

import (
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/whicencer/golang-auth/database"
	"github.com/whicencer/golang-auth/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

const SecretKey = "secret"

func Register(c *fiber.Ctx) error {
	// init a "db" and "user" variable
	db := database.DB.Db

	var body struct {
		Nickname    string
		Description string
		Password    string
	}

	// Getting an error on parsing a request body data
	// If there is an error - return a StatusBadRequest with specified JSON
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	// Define a hashed password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	// If there are some errors by hashing a password - return a Status Internal Error with JSON
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to hash password",
		})
	}

	user := models.User{Nickname: body.Nickname, Description: body.Description, Password: string(hashedPassword)}

	if err := db.Create(&user).Error; err != nil {
		if strings.Contains(err.Error(), "duplicated key not allowed") {
			return c.Status(fiber.StatusConflict).JSON(fiber.Map{
				"message": "Nickname already in use",
			})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to create user",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User created",
		"user":    user,
	})
}

func Login(c *fiber.Ctx) error {
	db := database.DB.Db
	var body struct {
		Nickname string `json:"nickname"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	var dbUser models.User
	if err := db.Where("nickname = ?", body.Nickname).First(&dbUser).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Invalid login or password!",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to retrieve user",
		})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(body.Password)); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid login or password!",
		})
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(dbUser.ID)),
		ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
	})

	token, err := claims.SignedString([]byte(SecretKey))

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error occured",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Successfully logged in",
		"token":   token,
	})
}
