package env

import (
	"fmt"
	"os"

	"github.com/pterm/pterm"
)

var (
	SchedulerSetKey      string
	PropagateUrl         string
	PropagateHeaderKey   string
	PropagateHeaderValue string
	PropagateMethod      string
)

func loadSchedulerEnv() {
	SchedulerSetKey = os.Getenv("SCHEDULER_SET_KEY")
	PropagateUrl = os.Getenv("PROPAGATE_URL")
	PropagateHeaderKey = os.Getenv("PROPAGATE_HEADER_KEY")
	PropagateHeaderValue = os.Getenv("PROPAGATE_HEADER_VALUE")
	PropagateMethod = os.Getenv("PROPAGATE_METHOD")

	pterm.DefaultLogger.Info(
		fmt.Sprintf("Scheduler environment done with set key %s", SchedulerSetKey),
	)
}
