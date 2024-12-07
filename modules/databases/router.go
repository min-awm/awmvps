package databases

import (
	"github.com/gofiber/fiber/v2"
)

func Init(app *fiber.App) {
	api := app.Group("/databases")
	api.Get("/install-mariadb", InstallMariaDB)
	api.Get("/status-mariadb", StatusMariaDB)
	api.Get("/start-mariadb", StartMariaDB)
	api.Get("/stop-mariadb", StopMariaDB)
	api.Get("/list-user-mariadb", ListUserMariaDB)
	api.Post("/create-user-mariadb", CreateUserMariaDB)
	api.Post("/delete-user-mariadb", DeleteUserMariaDB)
	api.Get("/remove-mariadb", RemoveMariaDB)
}
