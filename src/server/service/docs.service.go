package server_service

import (
	"fmt"

	_ "github.com/Rfluid/http-scheduler/docs"
	"github.com/Rfluid/http-scheduler/src/config/env"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"github.com/pterm/pterm"
)

func makeDocs(app *fiber.App) {
	app.Get("/docs/*", swagger.HandlerDefault)
	pterm.DefaultLogger.Info(
		fmt.Sprintf("Docs available at %s:%s/docs", env.ServerHost, env.ServerPort),
	)
}
