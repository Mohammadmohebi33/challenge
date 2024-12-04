package userhandler

import (
	"github.com/gofiber/fiber/v2"
	"hotel_with_test/params"
)

func (h UserHandler) Profile(c *fiber.Ctx) error {
	var req params.ProfileRequest
	if err := c.BodyParser(&req); err != nil {
		return err
	}
	user, err := h.UserService.Profile(c.Context(), req)
	if err != nil {
		return err
	}
	return c.JSON(user)
}
