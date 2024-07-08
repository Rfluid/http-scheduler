package env

import (
	"fmt"
	"os"
	"strconv"

	"github.com/pterm/pterm"
)

var (
	RedisAddr     string
	RedisPassword string
	RedisDB       int
)

func loadRedisEnv() {
	RedisAddr = os.Getenv("REDIS_ADDR")
	RedisPassword = os.Getenv("REDIS_PASSWORD")
	RedisDBStr := os.Getenv("REDIS_DB")

	var err error = nil
	RedisDB, err = strconv.Atoi(RedisDBStr)
	if err != nil {
		pterm.DefaultLogger.Error("Failed to convert REDIS_DB to int")
		os.Exit(1)
	}

	pterm.DefaultLogger.Info(
		fmt.Sprintf("Redis environment done with host %s", RedisAddr),
	)
}
