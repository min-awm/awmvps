package database

import (
	"github.com/gofiber/storage/bbolt/v2"
)

var Storage *bbolt.Storage

func Init() {
	Storage = bbolt.New(bbolt.Config{
		Database: "awmvps.db",
		Bucket:   "awmvps_bucket",
	})
}

func Close() {
	if Storage != nil {
		Storage.Close()
	}
}
