package databases

import (
	"awmvps/database"
	"awmvps/utils"
	"bytes"
	"os/exec"

	"github.com/gofiber/fiber/v2"
)

func InstallMariaDB(c *fiber.Ctx) error {
	osType := utils.DetectOS()

	if osType == "deb" {
		if output, err := utils.RunCommand("apt", "install", "-y", "mariadb-server", "mariadb-client"); err != nil {
			return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": false, "message": output})
		}
		if output, err := utils.RunCommand("systemctl", "enable", "mariadb"); err != nil {
			return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": false, "message": output})
		}
		if output, err := utils.RunCommand("systemctl", "start", "mariadb"); err != nil {
			return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": false, "message": output})
		}

		dbRootPassword := utils.RandomString(20)
		database.Storage.Set("dbRootPassword", []byte(dbRootPassword), 0)

		// Lệnh SQL cần thực thi
		sqlCommands := `
			ALTER USER 'root'@'localhost' IDENTIFIED BY '` + dbRootPassword + `';
			DELETE FROM mysql.user WHERE User='';
			DELETE FROM mysql.user WHERE User='root' AND Host NOT IN ('localhost', '127.0.0.1', '::1');
			DROP DATABASE IF EXISTS test;
			DELETE FROM mysql.db WHERE Db='test' OR Db='test\\_%';
			FLUSH PRIVILEGES;
		`

		// Tạo lệnh thực thi MySQL
		cmd := exec.Command("mysql", "-uroot")
		cmd.Stdin = bytes.NewBufferString(sqlCommands)

		output, err := cmd.CombinedOutput()
		if err != nil {
			return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": false, "message": output})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": true, "message": "Cài đặt thành công MariaDB"})
	} else if osType == "rpm" {
		if output, err := utils.RunCommand("yum", "install", "-y", "mariadb-server", "mariadb"); err != nil {
			return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": false, "message": output})
		}
		if output, err := utils.RunCommand("systemctl", "enable", "mariadb"); err != nil {
			return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": false, "message": output})
		}
		if output, err := utils.RunCommand("systemctl", "start", "mariadb"); err != nil {
			return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": false, "message": output})
		}

		dbRootPassword := utils.RandomString(20)
		database.Storage.Set("dbRootPassword", []byte(dbRootPassword), 0)

		// Lệnh SQL cần thực thi
		sqlCommands := `
			ALTER USER 'root'@'localhost' IDENTIFIED BY '` + dbRootPassword + `';
			DELETE FROM mysql.user WHERE User='';
			DELETE FROM mysql.user WHERE User='root' AND Host NOT IN ('localhost', '127.0.0.1', '::1');
			DROP DATABASE IF EXISTS test;
			DELETE FROM mysql.db WHERE Db='test' OR Db='test\\_%';
			FLUSH PRIVILEGES;
		`

		// Tạo lệnh thực thi MySQL
		cmd := exec.Command("mysql", "-uroot")
		cmd.Stdin = bytes.NewBufferString(sqlCommands)

		output, err := cmd.CombinedOutput()
		if err != nil {
			return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": false, "message": string(output)})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": true, "message": "Cài đặt thành công MariaDB"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": false, "message": "Không hỗ trợ"})
}

func StatusMariaDB(c *fiber.Ctx) error {
	_, err := utils.RunCommand("which", "mariadb")
	if err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": false, "message": "notInstall"})
	}

	version, err := utils.RunCommand("mysql", "--version")
	if err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": false, "message": "versionErr"})
	}

	status, _ := utils.RunCommand("systemctl", "is-active", "mariadb")

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": true, "version": version, "status": status})
}

func StartMariaDB(c *fiber.Ctx) error {
	if output, err := utils.RunCommand("systemctl", "start", "mariadb"); err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": false, "message": output})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": true, "message": "Đã bật MariaDB"})
}

func StopMariaDB(c *fiber.Ctx) error {
	if output, err := utils.RunCommand("systemctl", "stop", "mariadb"); err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": false, "message": output})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": true, "message": "Đã dừng MariaDB"})
}

func ListUserMariaDB(c *fiber.Ctx) error {
	sqlCommands := `
		SELECT User, Host
		FROM mysql.user
		WHERE User LIKE 'awmvps_%';
	`
	dbRootPassword, _ := database.Storage.Get("dbRootPassword")

	cmd := exec.Command("mysql", "-uroot", "-p"+string(dbRootPassword))
	cmd.Stdin = bytes.NewBufferString(sqlCommands)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": false, "message": string(output)})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": true, "message": string(output)})
}

func CreateUserMariaDB(c *fiber.Ctx) error {
	newUser := c.FormValue("newUser")
	ip := c.FormValue("ip")
	password := c.FormValue("password")
	sqlCommands := `
		CREATE USER '` + newUser + `'@'` + ip + `' IDENTIFIED BY '` + password + `';
		GRANT ALL PRIVILEGES ON *.* TO '` + newUser + `'@'` + ip + `' WITH GRANT OPTION;
		FLUSH PRIVILEGES;
	`
	dbRootPassword, _ := database.Storage.Get("dbRootPassword")

	cmd := exec.Command("mysql", "-uroot", "-p"+string(dbRootPassword))
	cmd.Stdin = bytes.NewBufferString(sqlCommands)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": false, "message": string(output)})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": true, "message": string(output)})
}

func DeleteUserMariaDB(c *fiber.Ctx) error {
	newUser := c.FormValue("newUser")
	ip := c.FormValue("ip")
	sqlCommands := `
		DROP USER IF EXISTS '` + newUser + `'@'` + ip + `';
	`
	dbRootPassword, _ := database.Storage.Get("dbRootPassword")

	cmd := exec.Command("mysql", "-uroot", "-p"+string(dbRootPassword))
	cmd.Stdin = bytes.NewBufferString(sqlCommands)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": false, "message": string(output)})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": true, "message": string(output)})
}

func RemoveMariaDB(c *fiber.Ctx) error {
	osType := utils.DetectOS()

	if osType == "deb" {
		if output, err := utils.RunCommand("systemctl", "stop", "mariadb"); err != nil {
			return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": false, "message": output})
		}

		if output, err := utils.RunCommand("apt", "remove", "--purge", "-y", "mariadb-server", "mariadb-client", "mariadb-common"); err != nil {
			return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": false, "message": output})
		}

		if output, err := utils.RunCommand("rm", "-rf", "/etc/mysql", "/etc/my.cnf /etc/my.cnf.d", "/var/lib/mysql", "/var/log/mysql", "/var/log/mariadb"); err != nil {
			return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": false, "message": output})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": true, "message": "Gỡ cài đặt thành công MariaDB"})
	} else if osType == "rpm" {
		if output, err := utils.RunCommand("systemctl", "stop", "mariadb"); err != nil {
			return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": false, "message": output})
		}

		if output, err := utils.RunCommand("yum", "remove", "-y", "mariadb", "mariadb-server", "mariadb-libs", "mariadb-common"); err != nil {
			return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": false, "message": output})
		}

		if output, err := utils.RunCommand("rm", "-rf", "/etc/mysql", "/etc/my.cnf /etc/my.cnf.d", "/var/lib/mysql", "/var/log/mysql", "/var/log/mariadb"); err != nil {
			return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": false, "message": output})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": true, "message": "Gỡ cài đặt thành công MariaDB"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": false, "message": "Không hỗ trợ"})
}
