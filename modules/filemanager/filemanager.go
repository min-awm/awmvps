package filemanager

import (
	"awmvps/utils"
	"os"
	"os/exec"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func CheckPath(c *fiber.Ctx) error {
	queryPath := c.Query("path", os.Getenv("HOME"))

	if utils.IsPath(queryPath) {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": true, "path": queryPath})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": false, "path": queryPath})
}

func List(c *fiber.Ctx) error {
	queryPath := c.Query("path", os.Getenv("HOME"))
	cmd := exec.Command("bash", "-c", "cd "+queryPath+" && ls -F")

	// Lấy kết quả
	output, err := cmd.Output()
	if err != nil {
		log.Error(err)
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": false, "message": "cmdErr"})
	}

	lines := strings.Split(string(output), "\n")
	results := []map[string]string{}

	// {name, type}
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		entryType := "file"
		name := line

		if strings.HasSuffix(line, "*") {
			entryType = "application"
			name = strings.TrimSuffix(line, "*")
		} else if strings.HasSuffix(line, "/") {
			entryType = "directory"
			name = strings.TrimSuffix(line, "/")
		}

		results = append(results, map[string]string{
			"name": name,
			"type": entryType,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": true, "list": results})
}

func Rename(c *fiber.Ctx) error {
	oldName := c.FormValue("oldName")
	newName := c.FormValue("newName")

	cmd := exec.Command("bash", "-c", "mv "+oldName+" "+newName)

	// Lấy kết quả
	_, err := cmd.Output()
	if err != nil {
		log.Error(err)
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": false, "message": "cmdErr"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": true, "message": ""})
}
