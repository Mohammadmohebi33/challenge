package middleware

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"hotel_with_test/repository/mongo/userrepo"
	"hotel_with_test/service/authservice"
	"strings"
	"time"
)

func Auth(auth authservice.Service, userStore *userrepo.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		headers, ok := c.GetReqHeaders()["Authorization"]
		if !ok {
			return fiber.ErrUnauthorized
		}
		token := strings.Replace(headers[0], "Bearer ", "", 1)

		fmt.Println(token)
		clams, err := auth.ParseToken(token)
		if err != nil {
			return err
		}
		expireFloat := clams["exp"].(float64)
		expire := int64(expireFloat)
		if time.Now().Unix() > expire {
			return errors.New("token expired")
		}
		userID := clams["id"].(string)
		user, err := userStore.GetUserByID(c.Context(), userID)
		if err != nil {
			return errors.New("user not found")
		}
		c.Context().SetUserValue("user", user)
		return c.Next()
	}
}
