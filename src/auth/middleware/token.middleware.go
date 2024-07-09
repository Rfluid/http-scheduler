package auth_middleware

import (
	"fmt"

	"github.com/Rfluid/http-scheduler/src/config/env"
	"github.com/gofiber/fiber/v2"
)

func TokenMiddleware(c *fiber.Ctx) error {
	if env.AuthToken == "" {
		return c.Next()
	}

	authHeader := c.Get(env.AuthHeaderKey)
	if authHeader == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": fmt.Sprintf("%s header is missing", env.AuthHeaderKey)})
	}
	if authHeader != env.AuthToken {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Invalid token"})
	}

	return c.Next()
}
