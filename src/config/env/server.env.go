package env

import (
	"fmt"
	"os"

	"github.com/pterm/pterm"
)

var (
	ServerHost string
	ServerPort string
)

func loadServerEnv() {
	ServerHost = os.Getenv("HOST")
	ServerPort = os.Getenv("PORT")

	pterm.DefaultLogger.Info(
		fmt.Sprintf("Server environment done with port %s", ServerPort),
	)
}
