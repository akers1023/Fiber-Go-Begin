package middleware

import (
	"fmt"
	"net/http"

	"github.com/akers1023/utils"
	"github.com/gofiber/fiber/v2"
)

func Authenticate() fiber.Handler {
	return func(c *fiber.Ctx) error {
		clientToken := c.Get("token")
		if clientToken == "" {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": fmt.Sprintf("No Authorization header provided")})
		}

		claims, err := utils.ValidateToken(clientToken)
		if err != "" {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err})
		}

		c.Locals("email", claims.Email)
		c.Locals("uid", claims.Uid)
		return c.Next()
	}
}
