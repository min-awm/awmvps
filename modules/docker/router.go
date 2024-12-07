package docker

import (
	"github.com/gofiber/fiber/v2"
)

func Init(app *fiber.App) {
	api := app.Group("/docker")
	api.Get("/status", Status)
	api.Get("/list-container", ListContainer)
	api.Post("/stop-container", StopContainer)
	api.Post("/start-container", StartContainer)
	api.Post("/stop-container", StopContainer)
	api.Post("/remove-container", RemoveContainer)

	api.Get("/list-image", ListImage)
	api.Post("/remove-image", RemoveImage)

	api.Get("/list-volume", ListVolume)
	api.Post("/remove-volume", RemoveVolume)
}
