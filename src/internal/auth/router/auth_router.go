package router

import (
	authhandler "golang-codebase/src/internal/auth/handler"

	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App, authHandler *authhandler.Handler) {
	api := app.Group("/api")
	v1 := api.Group("/v1")

	v1.Post("/login", authHandler.PostLogin)
	v1.Post("/registration", authHandler.PostRegistration)
}
