package security

import (
	"awmvps/utils"
	"os/exec"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func ListBlockIp(c *fiber.Ctx) error {
	output, err := utils.RunCommand("bash", "-c", "iptables -L INPUT -v -n | grep DROP | awk '{print $8}'")
	if err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": false, "message": output})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": true, "message": output})
}

func AddBlockIp(c *fiber.Ctx) error {
	ip := c.FormValue("ip")
	output, err := utils.RunCommand("iptables", "-A", "INPUT", "-s", ip, "-j", "DROP")
	if err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": false, "message": output})
	}

	err = utils.SaveIptablesConfig()
	if err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": false, "message": err})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": true, "message": output})
}

func RemoveBlockIp(c *fiber.Ctx) error {
	ip := c.FormValue("ip")
	output, err := utils.RunCommand("iptables", "-D", "INPUT", "-s", ip, "-j", "DROP")
	if err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": false, "message": output})
	}

	err = utils.SaveIptablesConfig()
	if err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": false, "message": err})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": true, "message": output})
}

func ListPort(c *fiber.Ctx) error {
	output, err := utils.RunCommand("bash", "-c", "iptables -L -v -n --line-numbers | awk '/dpt:/ {print $11,$12,$13}'")
	if err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": false, "message": output})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": true, "message": output})
}

func OpenPort(c *fiber.Ctx) error {
	port := c.FormValue("port")
	protocol := c.FormValue("protocol")
	output, err := utils.RunCommand("iptables", "-A", "INPUT", "-p", protocol, "--dport", port, "-j", "ACCEPT")
	if err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": false, "message": output})
	}

	err = utils.SaveIptablesConfig()
	if err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": false, "message": err})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": true, "message": output})
}

func DropPort(c *fiber.Ctx) error {
	port := c.FormValue("port")
	protocol := c.FormValue("protocol")
	output, err := utils.RunCommand("iptables", "-D", "INPUT", "-p", protocol, "--dport", port, "-j", "ACCEPT")
	if err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": false, "message": output})
	}

	err = utils.SaveIptablesConfig()
	if err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": false, "message": err})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": true, "message": output})
}

func DropAllPort(c *fiber.Ctx) error {
	if output, err := utils.RunCommand("iptables", "-P", "INPUT", "DROP"); err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": false, "message": output})
	}

	if output, err := utils.RunCommand("iptables", "-A", "INPUT", "-p", "tcp", "--dport", "ssh", "-j", "ACCEPT"); err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": false, "message": output})
	}

	if output, err := utils.RunCommand("iptables", "-A", "INPUT", "-p", "tcp", "--dport", "3000", "-j", "ACCEPT"); err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": false, "message": output})
	}

	if output, err := utils.RunCommand("iptables", "-A", "INPUT", "-i", "lo", "-j", "ACCEPT"); err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": false, "message": output})
	}

	err := utils.SaveIptablesConfig()
	if err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": false, "message": err})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": true})
}

func UndoDropAllPort(c *fiber.Ctx) error {
	if output, err := utils.RunCommand("iptables", "-P", "INPUT", "ACCEPT"); err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": false, "message": output})
	}

	err := utils.SaveIptablesConfig()
	if err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": false, "message": err})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": true})
}

func AllRole(c *fiber.Ctx) error {
	role := c.FormValue("role")

	if output, err := utils.RunCommand("bash", "-c", role); err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": false, "message": output})
	}

	err := utils.SaveIptablesConfig()
	if err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": false, "message": err})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": true})
}

func IptablesRestore() {
	if !utils.FileExists("iptables_backup.conf") {
		return
	}

	cmd := exec.Command("iptables-restore", "iptables_backup.conf")

	// Chạy lệnh
	if err := cmd.Run(); err != nil {
		log.Error("Failed to restore iptables config: ", err)
	}
}
