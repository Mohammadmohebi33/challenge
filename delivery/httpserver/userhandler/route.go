package userhandler

import "github.com/gofiber/fiber/v2"

func (h UserHandler) SetRoutes(e *fiber.App) {
	userGroup := e.Group("/users")

	userGroup.Post("/register", h.Register)
}
