package roomhandler

import "github.com/gofiber/fiber/v2"

func (h RoomHandler) GetAll(c *fiber.Ctx) error {
	rooms, err := h.roomService.GetAllRooms(c.Context())
	if err != nil {
		return err
	}
	return c.JSON(rooms)
}
