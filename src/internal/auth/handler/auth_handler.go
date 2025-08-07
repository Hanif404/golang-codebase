package authhandler

import (
	"golang-codebase/src/configs"
	auth "golang-codebase/src/domain/auth"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	service auth.Service
}

func New(s auth.Service) *Handler {
	return &Handler{
		service: s,
	}
}

func (h *Handler) PostLogin(c *fiber.Ctx) error {
	// Get the request body
	req := new(auth.Login)
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	errValid := configs.ValidateStruct(req)
	if errValid != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"validation_errors": errValid,
		})
	}

	// Call the service method
	resp, err := h.service.Login(req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal server error",
			"error":   err.Error(),
		})
	}

	// Return the response
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success",
		"data": fiber.Map{
			"token": resp,
		},
	})
}

func (h *Handler) PostRegistration(c *fiber.Ctx) error {
	// Get the request body
	req := new(auth.Registration)
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	errValid := configs.ValidateStruct(req)
	if errValid != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"validation_errors": errValid,
		})
	}

	// Call the service method
	err := h.service.Registration(req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal server error",
			"error":   err.Error(),
		})
	}

	// Return the response
	return c.Status(fiber.StatusOK).JSON("success")
}
