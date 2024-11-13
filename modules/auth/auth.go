package auth

import (
	"awmvps/database"
	"awmvps/utils"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/golang-jwt/jwt/v5"
)

func Login(c *fiber.Ctx) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	usernameAdmin, err := database.Storage.Get("usernameAdmin")
	if err != nil {
		return c.JSON(fiber.Map{"success": false, "message": "userPassIncorrect"})
	}

	passwordAdmin, err := database.Storage.Get("passwordAdmin")
	if err != nil {
		return c.JSON(fiber.Map{"success": false, "message": "userPassIncorrect"})
	}

	if username != string(usernameAdmin) || !utils.VerifyPassword(password, string(passwordAdmin)) {
		return c.JSON(fiber.Map{"success": false, "message": "userPassIncorrect"})
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"name": username,
		"exp":  time.Now().Add(time.Hour * 1).Unix(),
	})

	tokenKey, err := database.Storage.Get("tokenKey")
	if err != nil {
		log.Error(err)
	}

	t, err := token.SignedString(tokenKey)
	if err != nil {
		return c.JSON(fiber.Map{"success": false, "message": "serverError"})
	}

	return c.JSON(fiber.Map{"success": true, "token": t})
}
