package router

import (
	userhandler "golang-codebase/src/internal/user/handler"
	"golang-codebase/src/middleware"

	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App, userHandler *userhandler.Handler) {
	api := app.Group("/api")
	v1 := api.Group("/v1")

	v1.Get("/users", middleware.JWTAuth, userHandler.GetUsers)
	v1.Get("/user/:id", middleware.JWTAuth, userHandler.GetUserbyID)
	v1.Post("/user", middleware.JWTAuth, userHandler.PostUser)
	v1.Put("/user/:id", middleware.JWTAuth, userHandler.PutUser)
	v1.Delete("/user/:id", middleware.JWTAuth, userHandler.DeleteUser)
}
