package hotelhandler

import (
	"github.com/gofiber/fiber/v2"
	"hotel_with_test/delivery/httpserver/middleware"
)

func (h HotelHandler) SetRoute(e *fiber.App) {
	userGroup := e.Group("/hotel", middleware.Auth(h.authService, h.UserStore))

	userGroup.Post("/create", h.Create)
}
