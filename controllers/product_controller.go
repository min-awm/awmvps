package controllers

import (
	"awmvps/models"

	"github.com/gofiber/fiber/v2"
)

type ProductController struct{}

func (pc ProductController) GetProduct(c *fiber.Ctx) error {
	product := models.Product{
		ID:    1,
		Name:  "Laptop",
		Price: 1200.99,
	}
	return c.JSON(product)
}
