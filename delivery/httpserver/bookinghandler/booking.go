package bookinghandler

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"hotel_with_test/entity"
	"hotel_with_test/params"
)

func (h BookingHandler) BookRoom(c *fiber.Ctx) error {
	var bookParams params.BookRoomParams

	if err := c.BodyParser(&bookParams); err != nil {
		return err
	}

	roomID, err := primitive.ObjectIDFromHex(c.Params("roomID"))
	if err != nil {
		return err
	}

	user := c.Context().Value("user").(entity.MongoUser)

	resp, err := h.bookingService.Book(c.Context(), user.ID, roomID, bookParams)
	return c.JSON(resp)

}
