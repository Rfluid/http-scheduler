package scheduler

import (
	"github.com/Rfluid/http-scheduler/src/config/env"
	redis_scheduler "github.com/Rfluid/scheduler/src/redis"
	"github.com/pterm/pterm"
	"github.com/redis/go-redis/v9"

	local_redis "github.com/Rfluid/http-scheduler/src/redis"
)

var Worker *redis_scheduler.Worker

func LoadIntegration() {
	pterm.DefaultLogger.Info("Loading scheduler integration...")
	Worker = redis_scheduler.Create(local_redis.Connection(), env.SchedulerSetKey)
	Worker.SetCallback(func(data redis.Z) error {
		pterm.DefaultLogger.Info("propagate")
		return nil
	})
	pterm.DefaultLogger.Info("Scheduler integration loaded")
}
