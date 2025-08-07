package main

import (
	authhandler "golang-codebase/src/internal/auth/handler"
	authrepository "golang-codebase/src/internal/auth/repository"
	authRouter "golang-codebase/src/internal/auth/router"
	authservice "golang-codebase/src/internal/auth/service"
	userhandler "golang-codebase/src/internal/user/handler"
	userrepository "golang-codebase/src/internal/user/repository"
	userRouter "golang-codebase/src/internal/user/router"
	userservice "golang-codebase/src/internal/user/service"
	mysql "golang-codebase/src/pkg/database"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
)

func main() {
	app := fiber.New()
	app.Use(healthcheck.New())
	db := mysql.GetDatabase()

	userRepo := userrepository.New(db)
	userService := userservice.New(userRepo)
	userHandler := userhandler.New(userService)
	userRouter.Routes(app, userHandler)

	authRepo := authrepository.New(db)
	authService := authservice.New(authRepo)
	authHandler := authhandler.New(authService)
	authRouter.Routes(app, authHandler)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Listen(":3000")
}
