package server_service

import (
	"github.com/gofiber/fiber/v2"
)

func createStatusCheckEndpoint(app *fiber.App) {
	app.Get(
		"/status-check",
		statusCheckHandler,
	)
}

//	@Summary		Status check
//	@Description	Returns a 200 status if server is up
//	@Tags			Health
//	@Produce		json
//	@Success		200	{string}	string	"OK"
//	@Router			/status-check [get]
func statusCheckHandler(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusOK)
}
