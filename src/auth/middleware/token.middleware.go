package auth_middleware

import (
	"fmt"
	"strings"

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

	// Split the header to get the token
	splitToken := strings.Split(authHeader, "Bearer ")
	if len(splitToken) != 2 {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": fmt.Sprintf("Invalid %s header length", env.AuthHeaderKey)})
	}
	tokenString := splitToken[1]

	if tokenString != env.AuthToken {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Invalid token"})
	}

	return c.Next()
}
