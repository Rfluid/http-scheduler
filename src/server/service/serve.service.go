package server_service

import (
	"fmt"

	"github.com/Rfluid/http-scheduler/src/config/env"
	"github.com/gofiber/fiber/v2"
	"github.com/pterm/pterm"
)

func Serve() {
	app := fiber.New()

	makeDocs(app)
	createStatusCheckEndpoint(app)

	err := app.Listen(fmt.Sprintf(":%s", env.ServerPort))
	pterm.DefaultLogger.Fatal(
		fmt.Sprintf("%v", err),
	)
}
