package router

import (
	userhandler "golang-codebase/src/internal/user/handler"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App, userHandler *userhandler.Handler) {
	apiV1 := app.Group("/v1")

	apiV1.Get("/user/:id", userHandler.GetUserbyID)
	apiV1.Get("/users", userHandler.GetUsers)
	apiV1.Post("/user", userHandler.PostUser)
	apiV1.Put("/user/:id", userHandler.PutUser)
	apiV1.Delete("/user/:id", userHandler.DeleteUser)
}
