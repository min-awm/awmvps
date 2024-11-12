package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func VerifyPassword(password string, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

// b := utils.VerifyPassword("sssss", "$2a$10$RjDRtePXGB.t8OS1KsCIm.qlIvlKO2D6G3.nDG50ehrXIabXZ7sDi")
// fmt.Print(b)
