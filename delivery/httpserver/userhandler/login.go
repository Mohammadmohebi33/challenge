package userhandler

import (
	"github.com/gofiber/fiber/v2"
	"hotel_with_test/params"
)

func (h UserHandler) Login(c *fiber.Ctx) error {
	var req params.LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return err
	}

	resp, err := h.UserService.Login(req)
	if err != nil {
		return err
	}
	return c.JSON(resp)
}
