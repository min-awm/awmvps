package utils

import "os"

func IsPath(pathString string) bool {
	info, err := os.Stat(pathString)
	if os.IsNotExist(err) {
		return false
	}

	if err != nil {
		return false
	}

	return info.IsDir()
}
