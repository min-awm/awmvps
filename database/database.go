package database

import (
	"awmvps/utils"
	"crypto/rand"
	"encoding/hex"

	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/storage/bbolt/v2"
)

var Storage *bbolt.Storage

func Init() {
	Storage = bbolt.New(bbolt.Config{
		Database: "awmvps.db",
		Bucket:   "awmvps_bucket",
	})

	Storage.Set("tokenKey", []byte(generateTokenKey()), 0)

	usernameAdmin, err := Storage.Get("usernameAdmin")
	if err != nil {
		log.Error(err)
	}

	if usernameAdmin == nil {
		Storage.Set("usernameAdmin", []byte("admin"), 0)
		hashedPassword, err := utils.HashPassword("admin")
		if err != nil {
			log.Error(err)
		}
		Storage.Set("passwordAdmin", []byte(hashedPassword), 0)
	}
}

func Close() {
	if Storage != nil {
		Storage.Close()
	}
}

func generateTokenKey() string {
	secret := make([]byte, 32)
	_, err := rand.Read(secret)
	if err != nil {
		return "key"
	}
	return hex.EncodeToString(secret)
}
