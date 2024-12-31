package security

import (
	"awmvps/database"
	"awmvps/utils"
	"encoding/json"
	"os/exec"

	"github.com/coreos/go-iptables/iptables"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func getListIp() []string {
	list, _ := database.Storage.Get("listIp")
	var listArray []string
	err := json.Unmarshal(list, &listArray)
	if err != nil || listArray == nil {
		listArray = []string{}
	}

	return listArray
}

// Ip
func existIp(value string) bool {
	list := getListIp()
	for _, item := range list {
		if item == value {
			return true
		}
	}
	return false
}

func removeIp(list []string, target string) []string {
	var newList []string
	for _, item := range list {
		if item != target {
			newList = append(newList, item)
		}
	}
	return newList
}

func ListBlockIp(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": true, "message": getListIp()})
}

func AddBlockIp(c *fiber.Ctx) error {
	ip := c.FormValue("ip")

	exist := existIp(ip)
	if exist {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": false, "message": "IP already exists"})
	}

	ipt, _ := iptables.New()
	err := ipt.AppendUnique("filter", "INPUT", "-s", ip, "-j", "DROP")
	if err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": false, "message": err})
	}

	err = utils.SaveIptablesConfig()
	if err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": false, "message": err})
	}

	listIp := getListIp()
	listIp = append(listIp, ip)
	listIpBytes, err := json.Marshal(listIp)
	if err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": false, "message": "Error encoding IP list"})
	}

	database.Storage.Set("listIp", []byte(listIpBytes), 0)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": true})
}

func RemoveBlockIp(c *fiber.Ctx) error {
	ip := c.FormValue("ip")

	exist := existIp(ip)
	if !exist {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": false, "message": "Not found IP"})
	}

	ipt, _ := iptables.New()
	err := ipt.DeleteIfExists("filter", "INPUT", "-s", ip, "-j", "DROP")
	if err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": false, "message": err})
	}

	err = utils.SaveIptablesConfig()
	if err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": false, "message": err})
	}

	listIp := getListIp()
	listIp = removeIp(listIp, ip)
	listIpBytes, err := json.Marshal(listIp)
	if err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": false, "message": "Error encoding IP list"})
	}

	database.Storage.Set("listIp", []byte(listIpBytes), 0)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": true})
}

// Port
type PortType struct {
	Protocol string `json:"protocol"`
	Port     string `json:"port"`
	Status   string `json:"status"`
}

func getListPort() []PortType {
	list, _ := database.Storage.Get("listPort")
	var listArray []PortType
	err := json.Unmarshal(list, &listArray)
	if err != nil || listArray == nil {
		listArray = []PortType{}
	}

	return listArray
}

func existPort(port string, protocol string) (bool, PortType) {
	listPort := getListPort()
	for _, item := range listPort {
		if item.Port == port && item.Protocol == protocol {
			return true, item
		}
	}

	return false, PortType{}
}

func removePort(list []PortType, target PortType) []PortType {
	var newList []PortType
	for _, item := range list {
		if item != target {
			newList = append(newList, item)
		}
	}
	return newList
}

func ListPort(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": true, "message": getListPort()})
}

func AddPort(c *fiber.Ctx) error {
	portNumber := c.FormValue("port")
	protocol := c.FormValue("protocol")
	status := c.FormValue("status")

	ipt, _ := iptables.New()
	exist, portItem := existPort(portNumber, protocol)

	if exist && portItem.Status == status {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": false, "message": "Port already exists"})
	}

	if exist {
		// delete old port
		err := ipt.DeleteIfExists("filter", "INPUT", "-p", portItem.Protocol, "--dport", portItem.Port, "-j", portItem.Status)
		if err != nil {
			return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": false, "message": err})
		}

		listPort := getListPort()
		listPort = removePort(listPort, portItem)

		listPortBytes, err := json.Marshal(listPort)

		if err != nil {
			return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": false, "message": "Error encoding old port list"})
		}

		database.Storage.Set("listPort", []byte(listPortBytes), 0)
	}

	// add new port
	err := ipt.AppendUnique("filter", "INPUT", "-p", protocol, "--dport", portNumber, "-j", status)
	if err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": false, "message": err})
	}

	listPort := getListPort()
	listPort = append(listPort, PortType{
		Port:     portNumber,
		Protocol: protocol,
		Status:   status,
	})
	listPortBytes, err := json.Marshal(listPort)
	if err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": false, "message": "Error encoding new port list"})
	}

	database.Storage.Set("listPort", []byte(listPortBytes), 0)

	err = utils.SaveIptablesConfig()
	if err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": false, "message": err})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": true})
}

func RemovePort(c *fiber.Ctx) error {
	portNumber := c.FormValue("port")
	protocol := c.FormValue("protocol")
	status := c.FormValue("status")

	ipt, _ := iptables.New()
	err := ipt.DeleteIfExists("filter", "INPUT", "-p", protocol, "--dport", portNumber, "-j", status)
	if err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": false, "message": err})
	}

	listPort := getListPort()
	listPort = removePort(listPort, PortType{
		Port:     portNumber,
		Protocol: protocol,
		Status:   status,
	})

	listPortBytes, err := json.Marshal(listPort)

	if err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": false, "message": "Error encoding port list"})
	}

	database.Storage.Set("listPort", []byte(listPortBytes), 0)

	err = utils.SaveIptablesConfig()
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

	if err := cmd.Run(); err != nil {
		log.Error("Failed to restore iptables config: ", err)
	}
}
