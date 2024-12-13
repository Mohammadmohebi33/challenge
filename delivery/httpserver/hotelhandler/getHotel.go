package hotelhandler

import (
	"github.com/gofiber/fiber/v2"
)

func (h HotelHandler) GetHotel(c *fiber.Ctx) error {
	hotelID := c.Params("id")
	hotel, err := h.hotelService.GetHotelById(c.Context(), hotelID)
	if err != nil {
		return err
	}
	return c.JSON(hotel)
}

func (h HotelHandler) GetAllHotels(c *fiber.Ctx) error {
	hotels, err := h.hotelService.GetAllHotel(c.Context())
	if err != nil {
		return err
	}
	return c.JSON(hotels)
}

func (h HotelHandler) GetHotelsRoom(c *fiber.Ctx) error {
	hotelID := c.Params("hotelId")
	rooms, err := h.hotelService.GetRooms(c.Context(), hotelID)
	if err != nil {
		return err
	}
	return c.JSON(rooms)
}
