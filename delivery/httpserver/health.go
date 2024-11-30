package httpserver

import (
	"github.com/gofiber/fiber/v2"
)

func (s Server) healthCheck(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "everything is good!",
	})
}
