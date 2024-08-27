package worker_handler

import (
	"github.com/Rfluid/http-scheduler/src/integration/scheduler"
	"github.com/Rfluid/http-scheduler/src/redis"
	"github.com/gofiber/fiber/v2"
)

// @Summary		Tries to dequeue.
// @Description	Dequeues if the current date is after the date of the first element's score. Otherwise, schedules the dequeue.
// @Tags			Worker
// @Accept			json
// @Produce		json
// @Success		200		{string}	string						"OK"
// @Failure		400		{object}	fiber.Map					"Invalid request body"
// @Failure		500		{object}	fiber.Map					"Internal server error"
// @Router			/worker/try-dequeue [post]
func TryDequeue(c *fiber.Ctx) error {
	// Parse the request body
	err := scheduler.Worker.TryDequeue(redis.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	return c.SendStatus(fiber.StatusOK)
}
