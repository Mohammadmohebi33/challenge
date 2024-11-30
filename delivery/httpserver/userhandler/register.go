package userhandler

import (
	"github.com/gofiber/fiber/v2"
	"hotel_with_test/params"
)

func (h UserHandler) Register(c *fiber.Ctx) error {
	var req params.RegisterRequest
	if err := c.BodyParser(&req); err != nil {
		return err
	}

	user, err := h.UserService.Register(req)
	if err != nil {
		return err
	}

	return c.JSON(user)
}
