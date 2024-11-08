package main

import (
	"awm-vps/router"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// Thiết lập các router
	router.SetupRoutes(app)

	// Khởi chạy server
	app.Listen(":3000")
}
