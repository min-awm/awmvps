package main

import (
	"embed"
	"io/fs"
	"net/http"

	// "awmvps/modules/filemanager"
	"awmvps/database"
	"awmvps/modules/auth"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
)

//go:embed public/*
var embedDirPublic embed.FS

func main() {
	app := fiber.New()
	app.Use(cors.New())

	publicFiles, err := fs.Sub(embedDirPublic, "public")
	if err != nil {
		log.Error(err)
	}
	app.Use("/", filesystem.New(filesystem.Config{
		Root: http.FS(publicFiles),
	}))

	database.Init()
	defer database.Close()

	auth.Init(app)
	// filemanager.Init(app)

	app.Listen(":3000")
}
