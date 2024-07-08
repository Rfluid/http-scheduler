package server_service

import (
	"fmt"

	"github.com/Rfluid/http-scheduler/src/config/env"
	worker_router "github.com/Rfluid/http-scheduler/src/worker/router"
	"github.com/gofiber/fiber/v2"
	"github.com/pterm/pterm"
)

func Serve() {
	app := fiber.New()

	makeDocs(app)
	createStatusCheckEndpoint(app)

	worker_router.Route(app)

	err := app.Listen(fmt.Sprintf(":%s", env.ServerPort))
	pterm.DefaultLogger.Fatal(
		fmt.Sprintf("%v", err),
	)
}
