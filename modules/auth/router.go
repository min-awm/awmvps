package auth

import (
	"awmvps/database"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/golang-jwt/jwt/v5"
)

func Init(app *fiber.App) {
	api := app.Group("/auth")

	api.Post("/login", Login)
	api.Get("/", accessible)

	tokenKey, err := database.Storage.Get("tokenKey")
	if err != nil {
		log.Error(err)
	}

	app.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: tokenKey},
	}))

	api.Get("/restricted", restricted)
}

func accessible(c *fiber.Ctx) error {
	return c.SendString("Accessible")
}

func restricted(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	return c.SendString("Welcome " + name)
}
