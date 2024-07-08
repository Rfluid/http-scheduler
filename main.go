package main

import (
	"github.com/Rfluid/http-scheduler/src/config"
	"github.com/Rfluid/http-scheduler/src/integration"
	server_service "github.com/Rfluid/http-scheduler/src/server/service"
)

// @title						Http Scheduler
// @version					1.0
// @description				Schedules data propagations using HTTP requests
// @securityDefinitions.apikey	ApiKeyAuth
// @in							header
// @name						Authorization
func main() {
	config.Load()
	integration.Load()
	server_service.Serve()
}
