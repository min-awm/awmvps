package docker

import (
	"awmvps/utils"

	"github.com/gofiber/fiber/v2"
)

func Status(c *fiber.Ctx) error {
	_, err := utils.RunCommand("which", "docker")
	if err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": false, "message": "Chưa cài đặt docker"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": true, "message": "Đã cài đặt docker"})
}

func ListContainer(c *fiber.Ctx) error {
	output, err := utils.RunCommand("docker", "container", "ls", "-a", "--format", "'{{json .}}'")
	if err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": false, "message": output})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": true, "message": output})
}

func StopContainer(c *fiber.Ctx) error {
	id := c.FormValue("id")
	output, err := utils.RunCommand("docker", "container", "stop", id)
	if err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": false, "message": output})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": true, "message": "Đã dừng container " + id})
}

func StartContainer(c *fiber.Ctx) error {
	id := c.FormValue("id")
	output, err := utils.RunCommand("docker", "container", "start", id)
	if err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": false, "message": output})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": true, "message": "Đã chạy container " + id})
}

func RemoveContainer(c *fiber.Ctx) error {
	id := c.FormValue("id")
	output, err := utils.RunCommand("docker", "container", "rm", id)
	if err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": false, "message": output})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": true, "message": "Đã xóa container " + id})
}

func ListImage(c *fiber.Ctx) error {
	output, err := utils.RunCommand("docker", "image", "ls", "-a", "--format", "'{{json .}}'")
	if err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": false, "message": output})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": true, "message": output})
}

func RemoveImage(c *fiber.Ctx) error {
	id := c.FormValue("id")
	output, err := utils.RunCommand("docker", "image", "rm", id)
	if err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": false, "message": output})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": true, "message": "Đã xóa image " + id})
}

func ListVolume(c *fiber.Ctx) error {
	output, err := utils.RunCommand("docker", "volume", "ls", "--format", "'{{json .}}'")
	if err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": false, "message": output})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": true, "message": output})
}

func RemoveVolume(c *fiber.Ctx) error {
	id := c.FormValue("id")
	output, err := utils.RunCommand("docker", "volume", "rm", id)
	if err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": false, "message": output})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": true, "message": "Đã xóa volume " + id})
}
