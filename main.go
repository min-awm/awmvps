package main

import (
	// "awmvps/modules/filemanager"
	"awmvps/modules/auth"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	auth.Init(app)
	// filemanager.Init(app)

	app.Listen(":3000")
}
