package api

import (
	"fiberauthenticationjwt/entities"
	"fiberauthenticationjwt/services"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
	"github.com/spf13/viper"
)

type handler struct {
	service *services.Service
	app     *fiber.App
	secured fiber.Router
}

type errObj struct {
	ErrorCode int64  `json:"errorCode,omitempty"`
	ErrorMsg  string `json:"errorMsg,omitempty"`
}

//Handler api handler
func Handler(service *services.Service, app *fiber.App) {
	secured := app.Group("/app")
	appSecret := viper.Get("APP_SECRET_KEY").(string)
	if appSecret == "" {
		appSecret = "60ba4819aea76c0dc9470adbd48cd58f9a46ac92f07866ff590b08cdc2ae8a2d4035bc304bdedcaa7b47e379d700b2eb19246f006318c3ac85bcc182bded536f"
	}

	jwtConfig := jwtware.Config{
		Claims:        &entities.JwtApplicationUserClaim{},
		SigningKey:    []byte(appSecret),
		SigningMethod: "HS256",
	}
	secured.Use(jwtware.New(jwtConfig))

	handler := &handler{
		service: service,
		app:     app,
		secured: secured,
	}

	UserManagementHandler(handler)

}
