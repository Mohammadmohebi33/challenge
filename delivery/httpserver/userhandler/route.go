package userhandler

import (
	"github.com/gofiber/fiber/v2"
	"hotel_with_test/delivery/httpserver/middleware"
)

func (h UserHandler) SetRoutes(e *fiber.App) {
	userGroup := e.Group("/users")

	userGroup.Post("/register", h.Register)
	userGroup.Post("/login", h.Login)
	userGroup.Get("/profile", middleware.Auth(h.AuthService, h.UserStore), h.Profile)
}
