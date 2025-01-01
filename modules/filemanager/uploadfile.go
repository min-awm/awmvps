package filemanager

import (
	"path/filepath"

	"github.com/gofiber/fiber/v2"
)

func UploadFile(c *fiber.Ctx) error {
	file, err := c.FormFile("file")
	path := c.FormValue("path")
	if err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": false, "message": "Không thể lấy file từ yêu cầu"})
	}

	filePath := filepath.Join(path, file.Filename)
	if err := c.SaveFile(file, filePath); err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": false, "message": "Không thể lưu file"})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": true, "message": "File " + file.Filename + " đã được upload thành công!"})
}
