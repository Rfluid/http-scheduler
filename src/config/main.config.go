package config

import (
	"github.com/Rfluid/http-scheduler/src/config/env"
	redis_main "github.com/Rfluid/http-scheduler/src/redis/main"
)

func Load() {
	env.Load()
	redis_main.Redis()
}
