package env

import (
	"fmt"
	"os"

	"github.com/pterm/pterm"
)

var (
	SchedulerListKey string
	SchedulerLockKey string
)

func loadSchedulerEnv() {
	SchedulerListKey = os.Getenv("SCHEDULER_LIST_KEY")
	SchedulerLockKey = os.Getenv("SCHEDULER_LOCK_KEY")

	pterm.DefaultLogger.Info(
		fmt.Sprintf("Scheduler environment done with list key %s and lock key %s", SchedulerListKey, SchedulerLockKey),
	)
}
