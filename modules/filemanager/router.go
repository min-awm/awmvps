package filemanager

import (
	"github.com/gofiber/fiber/v2"
)

func Init(app *fiber.App) {
	api := app.Group("/filemanager")

	api.Get("/check-path", CheckPath)
	api.Get("/list", List)
	api.Post("/rename", Rename)
}
