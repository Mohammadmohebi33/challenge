package bookinghandler

import (
	"github.com/gofiber/fiber/v2"
	"hotel_with_test/delivery/httpserver/middleware"
)

func (h BookingHandler) SetRoute(e *fiber.App) {
	bookingGroup := e.Group("/booking", middleware.Auth(h.authService, h.UserStore))

	bookingGroup.Post("/book/:roomID", h.BookRoom)
}
