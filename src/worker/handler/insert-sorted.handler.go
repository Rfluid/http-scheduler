package worker_handler

import (
	worker_model "github.com/Rfluid/http-scheduler/src/worker/model"
	worker_service "github.com/Rfluid/http-scheduler/src/worker/service"
	"github.com/gofiber/fiber/v2"
)

// @Summary		Schedules callback at data.
// @Description	Inserts data to be executed in callback at specific date in scheduler worker.
// @Tags			Worker
// @Accept			json
// @Produce		json
// @Param			task	body		worker_model.InsertSorted	true	"Insert task"
// @Success		200		{string}	string						"OK"
// @Failure		400		{object}	fiber.Map					"Invalid request body"
// @Failure		500		{object}	fiber.Map					"Internal server error"
// @Router			/worker/insert-sorted [post]
func InsertSorted(c *fiber.Ctx) error {
	// Parse the request body
	var b worker_model.InsertSorted
	if err := c.BodyParser(&b); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid request body"})
	}

	err := worker_service.InsertSorted(b)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	return c.SendStatus(fiber.StatusOK)
}
