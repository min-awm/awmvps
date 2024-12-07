package utils

import (
	"os"
	"os/exec"
)

func DetectOS() string {
	if _, err := exec.LookPath("apt"); err == nil {
		return "deb"
	}

	if _, err := exec.LookPath("yum"); err == nil {
		return "rpm"
	}

	return "notSupport"
}

func RunCommand(command string, args ...string) (string, error) {
	cmd := exec.Command(command, args...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return string(output), err
	}
	return string(output), nil
}

func FileExists(filename string) bool {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return err == nil
}

func SaveIptablesConfig() error {
	file, err := os.Create("iptables_backup.conf")
	if err != nil {
		return err
	}
	defer file.Close()

	cmd := exec.Command("iptables-save")
	cmd.Stdout = file

	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}
