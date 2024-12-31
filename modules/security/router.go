package security

import (
	"github.com/gofiber/fiber/v2"
)

func Init(app *fiber.App) {
	IptablesRestore()
	api := app.Group("/security")

	api.Get("/list-block-ip", ListBlockIp)
	api.Post("/add-block-ip", AddBlockIp)
	api.Post("/remove-block-ip", RemoveBlockIp)

	api.Get("/list-port", ListPort)
	api.Post("/add-port", AddPort)
	api.Post("/remove-port", RemovePort)

	api.Post("/add-role", AllRole)
}
