package nginx

import (
	"awmvps/utils"

	"github.com/gofiber/fiber/v2"
)

func Install(c *fiber.Ctx) error {
	osType := utils.DetectOS()

	if osType == "deb" {
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
	} else if osType == "rpm" {
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
	osType := utils.DetectOS()

	if osType == "deb" {
		if output, err := utils.RunCommand("systemctl", "stop", "nginx"); err != nil {
			return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": false, "message": output})
		}

		if output, err := utils.RunCommand("apt", "remove", "nginx"); err != nil {
			return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": false, "message": output})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": true, "message": "Gỡ cài đặt thành công Nginx"})
	} else if osType == "rpm" {
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
	output, err := utils.RunCommand("bash", "-c", "cd /etc/nginx/conf.d && ls *.conf")
	if err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": false, "message": output})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": true, "message": output})
}

func GetFileConf(c *fiber.Ctx) error {
	fileName := c.FormValue("fileName")
	output, err := utils.RunCommand("cat", "/etc/nginx/conf.d/"+fileName)
	if err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": false, "message": output})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": true, "message": output})
}

func CreateConf(c *fiber.Ctx) error {
	fileName := c.FormValue("fileName")
	if output, err := utils.RunCommand("sh", "-c", `echo -e "server {
	listen 80;
	server_name example.com;
	# Replace with your server_name

	location / {
		# Replace with your app
		proxy_pass http://localhost:3000;
		proxy_http_version 1.1;
		proxy_set_header Upgrade \$http_upgrade;
		proxy_set_header Connection 'upgrade';
		proxy_set_header Host \$host;
		proxy_cache_bypass \$http_upgrade;
	}
}" | tee /etc/nginx/conf.d/`+fileName+` > /dev/null`); err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": false, "message": output})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": true, "message": "Đã tạo file: " + fileName})
}

func RemoveConf(c *fiber.Ctx) error {
	fileName := c.FormValue("fileName")
	if output, err := utils.RunCommand("rm", "/etc/nginx/conf.d/"+fileName); err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": false, "message": output})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": true, "message": "Đã xóa file: " + fileName})
}

func SaveConf(c *fiber.Ctx) error {
	fileName := c.FormValue("fileName")
	content := c.FormValue("content")
	if output, err := utils.RunCommand("sh", "-c", `echo -e "`+content+`" | tee /etc/nginx/conf.d/`+fileName+` > /dev/null`); err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": false, "message": output})
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
