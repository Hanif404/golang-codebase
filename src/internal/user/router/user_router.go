package router

import (
	userhandler "golang-codebase/src/internal/user/handler"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App, userHandler *userhandler.Handler) {
	api := app.Group("/api")
	v1 := api.Group("/v1")

	v1.Get("/users", userHandler.GetUsers)
	v1.Get("/user/:id", userHandler.GetUserbyID)
	v1.Post("/user", userHandler.PostUser)
	v1.Put("/user/:id", userHandler.PutUser)
	v1.Delete("/user/:id", userHandler.DeleteUser)
}
