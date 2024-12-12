package auth

import (
	"awmvps/database"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func Init(app *fiber.App) {
	api := app.Group("/auth")

	api.Post("/login", Login)

	tokenKey, err := database.Storage.Get("tokenKey")
	if err != nil {
		log.Error(err)
	}

	app.Use(jwtware.New(jwtware.Config{
		SigningKey:  jwtware.SigningKey{Key: tokenKey},
		TokenLookup: "header:Authorization,query:token",
		AuthScheme:  "Bearer",
	}))

	api.Get("/user-info", UserInfo)
	api.Post("/change-user-info", ChangeUserInfo)
	api.Get("/package-manager", GetPackageManager)
}
