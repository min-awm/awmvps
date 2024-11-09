package controllers

import (
	"awmvps/models"

	"github.com/gofiber/fiber/v2"
)

type UserController struct{}

func (uc UserController) GetUser(c *fiber.Ctx) error {
	user := models.User{
		ID:    1,
		Name:  "John Doe",
		Email: "john.doe@example.com",
	}
	return c.JSON(user)
}
