package nginx

import (
	"github.com/gofiber/fiber/v2"
)

func Init(app *fiber.App) {
	api := app.Group("/nginx")
	api.Get("/install", Install)
	api.Get("/status", Status)
	api.Get("/start", Start)
	api.Get("/stop", Stop)
	api.Get("/remove", Remove)
	api.Get("/list-conf", ListConf)
	api.Post("/file-conf", GetFileConf)
	api.Post("/create-conf", CreateConf)
	api.Post("/remove-conf", RemoveConf)
	api.Post("/save-conf", SaveConf)
	api.Get("/check-conf", CheckConf)
}
