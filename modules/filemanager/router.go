package filemanager

import (
	"github.com/gofiber/fiber/v2"
)

func Init(app *fiber.App) {
	api := app.Group("/filemanager")

	api.Get("/user/:id")
}
