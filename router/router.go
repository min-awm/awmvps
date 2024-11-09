package router

import (
	"awmvps/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	userController := controllers.UserController{}
	productController := controllers.ProductController{}

	// Route cho người dùng (User)
	app.Get("/user/:id", userController.GetUser)

	// Route cho sản phẩm (Product)
	app.Get("/product/:id", productController.GetProduct)
}
