package main

import (
	userhandler "golang-codebase/src/internal/user/handler"
	userrepository "golang-codebase/src/internal/user/repository"
	"golang-codebase/src/internal/user/router"
	userservice "golang-codebase/src/internal/user/service"
	mysql "golang-codebase/src/pkg/database"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	db := mysql.GetDatabase()

	userRepo := userrepository.New(db)
	userService := userservice.New(userRepo)
	userHandler := userhandler.New(userService)

	router.UserRoutes(app, userHandler)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Listen(":3000")
}
