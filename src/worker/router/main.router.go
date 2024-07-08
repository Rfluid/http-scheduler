package worker_router

import (
	worker_handler "github.com/Rfluid/http-scheduler/src/worker/handler"
	"github.com/gofiber/fiber/v2"
)

func Route(app *fiber.App) {
	group := app.Group("/worker")

	mainRoutes(group)
}

func mainRoutes(group fiber.Router) {
	group.Post("/insert-sorted", worker_handler.InsertSorted)
}
