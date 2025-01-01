package nginx

import (
	"awmvps/utils"
	"os"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
)

func Install(c *fiber.Ctx) error {
	packageManager := utils.DetectPackageManager()

	if packageManager == "apt" {
		if output, err := utils.RunCommand("apt", "install", "-y", "nginx"); err != nil {
			return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": false, "message": output})
		}
		if output, err := utils.RunCommand("systemctl", "enable", "nginx"); err != nil {
			return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": false, "message": output})
		}
		if output, err := utils.RunCommand("systemctl", "start", "nginx"); err != nil {
			return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": false, "message": output})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": true, "message": "Cài đặt thành công Nginx"})
	} else if packageManager == "yum" {
		if output, err := utils.RunCommand("yum", "install", "-y", "nginx"); err != nil {
			return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": false, "message": output})
		}
		if output, err := utils.RunCommand("systemctl", "enable", "nginx"); err != nil {
			return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": false, "message": output})
		}
		if output, err := utils.RunCommand("systemctl", "start", "nginx"); err != nil {
			return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": false, "message": output})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": true, "message": "Cài đặt thành công Nginx"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": false, "message": "Không hỗ trợ"})
}

func Status(c *fiber.Ctx) error {
	_, err := utils.RunCommand("which", "nginx")
	if err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": false, "message": "notInstall"})
	}

	status, _ := utils.RunCommand("systemctl", "is-active", "nginx")

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": true, "status": status})
}

func Start(c *fiber.Ctx) error {
	if output, err := utils.RunCommand("systemctl", "start", "nginx"); err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": false, "message": output})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": true, "message": "Đã bật Nginx"})
}

func Stop(c *fiber.Ctx) error {
	if output, err := utils.RunCommand("systemctl", "stop", "nginx"); err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": false, "message": output})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": true, "message": "Đã dừng Nginx"})
}

func Remove(c *fiber.Ctx) error {
	packageManager := utils.DetectPackageManager()

	if packageManager == "apt" {
		if output, err := utils.RunCommand("systemctl", "stop", "nginx"); err != nil {
			return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": false, "message": output})
		}

		if output, err := utils.RunCommand("apt", "remove", "nginx"); err != nil {
			return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": false, "message": output})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": true, "message": "Gỡ cài đặt thành công Nginx"})
	} else if packageManager == "yum" {
		if output, err := utils.RunCommand("systemctl", "stop", "nginx"); err != nil {
			return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": false, "message": output})
		}

		if output, err := utils.RunCommand("yum", "remove", "-y", "nginx"); err != nil {
			return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": false, "message": output})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": true, "message": "Gỡ cài đặt thành công Nginx"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": false, "message": "Không hỗ trợ"})
}

func ListConf(c *fiber.Ctx) error {
	dirPath := "/etc/nginx/conf.d/"

	files, err := filepath.Glob(filepath.Join(dirPath, "*.conf"))
	if err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": false, "message": "Lỗi khi tìm file"})
	}

	fileNames := []string{}
	for _, file := range files {
		fileNames = append(fileNames, filepath.Base(file))
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": true, "message": fileNames})
}

func GetFileConf(c *fiber.Ctx) error {
	fileName := c.FormValue("fileName")
	dirPath := "/etc/nginx/conf.d/"
	fullPath := filepath.Join(dirPath, fileName)

	content, err := os.ReadFile(fullPath)
	if err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": false, "message": "Lỗi khi đọc file " + fileName})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": true, "message": string(content)})
}

func CreateConf(c *fiber.Ctx) error {
	fileName := c.FormValue("fileName")
	content := `server {
	listen 80; # Lắng nghe trên cổng 80
	server_name example.com www.example.com; # Tên miền server

	root /var/www/example.com; # Thư mục gốc của website
	index index.html index.htm;

	location / {
		try_files $uri $uri/ =404; # Cách xử lý yêu cầu
	}

	error_page 404 /404.html; # Trang lỗi 404 tùy chỉnh
}`

	dirPath := "/etc/nginx/conf.d/"
	fullPath := filepath.Join(dirPath, fileName)

	file, err := os.Create(fullPath)
	if err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": false, "message": "Không thể tạo file " + fileName})

	}
	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": false, "message": "Không thể ghi vào file " + fileName})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": true, "message": "Đã tạo file: " + fileName})
}

func RemoveConf(c *fiber.Ctx) error {
	fileName := c.FormValue("fileName")
	dirPath := "/etc/nginx/conf.d/"
	fullPath := filepath.Join(dirPath, fileName)

	err := os.Remove(fullPath)
	if err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": false, "message": "Không thể xóa file " + fileName})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": true, "message": "Đã xóa file: " + fileName})
}

func SaveConf(c *fiber.Ctx) error {
	fileName := c.FormValue("fileName")
	content := c.FormValue("content")

	dirPath := "/etc/nginx/conf.d/"
	fullPath := filepath.Join(dirPath, fileName)

	file, err := os.Create(fullPath)
	if err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": false, "message": "Không thể lưu file " + fileName})

	}
	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": false, "message": "Không thể ghi vào file " + fileName})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": true, "message": "Đã lưu thay đổi: " + fileName})
}

func CheckConf(c *fiber.Ctx) error {
	output, err := utils.RunCommand("nginx", "-t")
	if err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": false, "message": output})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": true, "message": output})
}
