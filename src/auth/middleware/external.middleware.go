package auth_middleware

import (
	"fmt"
	"net/http"

	"github.com/Rfluid/http-scheduler/src/config/env"
	"github.com/gofiber/fiber/v2"
)

// Validate the request through http request.
func ExternalMiddleware(c *fiber.Ctx) error {
	if env.AuthEndpoint == "" {
		return c.Next()
	}

	// Get the authorization header
	authHeader := c.Get(env.AuthHeaderKey)
	if authHeader == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": fmt.Sprintf("%s header is missing", env.AuthHeaderKey)})
	}

	req, err := http.NewRequest(env.AuthMethod, env.AuthEndpoint, nil)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}
	req.Header = http.Header{env.AuthHeaderKey: {authHeader}}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": err.Error()})
	}
	if resp.StatusCode != http.StatusOK {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Unauthorized"})
	}

	return c.Next()
}
