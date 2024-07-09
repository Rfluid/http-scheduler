package env

import (
	"fmt"
	"os"

	"github.com/pterm/pterm"
)

var (
	AuthEndpoint  string
	AuthHeaderKey string
	AuthMethod    string
	AuthToken     string
)

func loadAuthEnv() {
	AuthEndpoint = os.Getenv("AUTH_ENDPOINT")
	AuthHeaderKey = os.Getenv("AUTH_HEADER_KEY")
	AuthMethod = os.Getenv("AUTH_METHOD")
	AuthToken = os.Getenv("AUTH_TOKEN")

	pterm.DefaultLogger.Info(
		fmt.Sprintf("Auth environment done with endpoint %s", AuthEndpoint),
	)
}
