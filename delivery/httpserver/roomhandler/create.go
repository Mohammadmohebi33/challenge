package roomhandler

import (
	"github.com/gofiber/fiber/v2"
	"hotel_with_test/params"
)

func (h RoomHandler) Create(c *fiber.Ctx) error {
	var req params.CreateRoomRequest

	if err := c.BodyParser(&req); err != nil {
		return fiber.ErrBadRequest
	}

	resp, err := h.roomService.CreateRoom(c.Context(), req)
	if err != nil {
		return fiber.ErrInternalServerError
	}
	return c.JSON(resp)
}
