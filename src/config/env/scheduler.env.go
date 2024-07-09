package env

import (
	"fmt"
	"os"

	"github.com/pterm/pterm"
)

var (
	SchedulerSetKey string
	PropagateUrl    string
)

func loadSchedulerEnv() {
	SchedulerSetKey = os.Getenv("SCHEDULER_SET_KEY")
	PropagateUrl = os.Getenv("PROPAGATE_URL")

	pterm.DefaultLogger.Info(
		fmt.Sprintf("Scheduler environment done with set key %s", SchedulerSetKey),
	)
}
