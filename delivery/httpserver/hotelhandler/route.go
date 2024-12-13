package hotelhandler

import (
	"github.com/gofiber/fiber/v2"
	"hotel_with_test/delivery/httpserver/middleware"
)

func (h HotelHandler) SetRoute(e *fiber.App) {
	hotelGroup := e.Group("/hotel", middleware.Auth(h.authService, h.UserStore))

	hotelGroup.Post("/create", h.Create)
	hotelGroup.Get("/get/:id", h.GetHotel)
	hotelGroup.Get("/get", h.GetAllHotels)
	hotelGroup.Get("/:hotelId/rooms", h.GetHotelsRoom)
}
