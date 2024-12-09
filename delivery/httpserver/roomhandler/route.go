package roomhandler

import (
	"github.com/gofiber/fiber/v2"
	"hotel_with_test/delivery/httpserver/middleware"
)

func (h RoomHandler) SetRoute(e *fiber.App) {
	roomGroup := e.Group("/room", middleware.Auth(h.authService, h.UserStore))

	roomGroup.Post("/create", h.Create)

}
