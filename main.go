package main

import (
	"embed"
	"io/fs"
	"net/http"

	"awmvps/database"
	"awmvps/modules/auth"
	"awmvps/modules/databases"
	"awmvps/modules/docker"
	"awmvps/modules/filemanager"
	"awmvps/modules/nginx"
	"awmvps/modules/security"
	"awmvps/modules/terminal"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
)

//go:embed public/*
var embedDirPublic embed.FS

func main() {
	app := fiber.New(fiber.Config{
		BodyLimit:             500 * 1024 * 1024,
		DisableStartupMessage: true,
	})
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
	terminal.Init(app)
	databases.Init(app)
	docker.Init(app)
	filemanager.Init(app)
	nginx.Init(app)
	security.Init(app)

	log.Info("My app is running at port 3000")
	app.Listen(":3000")

}
