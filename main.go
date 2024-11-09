package main

import (
	// "awmvps/modules/filemanager"
	"awmvps/database"
	"awmvps/modules/auth"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	database.Init()
	defer database.Close()

	auth.Init(app)
	// filemanager.Init(app)

	app.Listen(":3000")
}
