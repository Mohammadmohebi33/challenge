package hotelhandler

import (
	"github.com/gofiber/fiber/v2"
	"hotel_with_test/params"
)

func (h HotelHandler) Create(c *fiber.Ctx) error {
	var req params.HotelCreateRequest
	if err := c.BodyParser(&req); err != nil {
		return err
	}

	resp, err := h.hotelService.Create(c.Context(), req)
	if err != nil {
		return err
	}
	return c.JSON(resp)
}
