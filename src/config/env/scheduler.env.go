package env

import (
	"fmt"
	"os"

	"github.com/pterm/pterm"
)

var SchedulerSetKey string

func loadSchedulerEnv() {
	SchedulerSetKey = os.Getenv("SCHEDULER_SET_KEY")

	pterm.DefaultLogger.Info(
		fmt.Sprintf("Scheduler environment done with set key %s", SchedulerSetKey),
	)
}
