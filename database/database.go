package database

import (
	"awmvps/utils"

	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/storage/bbolt/v2"
)

var Storage *bbolt.Storage

func Init() {
	Storage = bbolt.New(bbolt.Config{
		Database: "awmvps.db",
		Bucket:   "awmvps_bucket",
	})

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
