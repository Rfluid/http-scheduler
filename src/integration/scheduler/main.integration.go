package scheduler

import (
	"github.com/Rfluid/http-scheduler/src/config/env"
	redis "github.com/Rfluid/http-scheduler/src/redis"
	redis_scheduler "github.com/Rfluid/scheduler/src/redis"
	"github.com/pterm/pterm"
)

var Worker *redis_scheduler.Worker

func LoadIntegration() {
	pterm.DefaultLogger.Info("Loading scheduler integration...")
	Worker = redis_scheduler.Create(redis.Connection(), env.SchedulerListKey, env.SchedulerLockKey)
	pterm.DefaultLogger.Info("Scheduler integration loaded")
}
