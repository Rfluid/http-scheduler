package redis_connect

import (
	"context"
	"fmt"
	"os"

	"github.com/Rfluid/http-scheduler/src/config/env"
	"github.com/pterm/pterm"
	"github.com/redis/go-redis/v9"
)

var (
	RedisClient *redis.ClusterClient
	Ctx         = context.Background() // Context passed to Redis client
)

func Connect() {
	pterm.DefaultLogger.Info("Connecting to redis...")
	RedisClient = redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: env.RedisAddrs,
	})

	// Test the connection
	_, err := RedisClient.Ping(Ctx).Result()
	if err != nil {
		pterm.DefaultLogger.Error(
			fmt.Sprintf("Unable to connect to Redis: %v", err),
		)
		os.Exit(1)
	}
	pterm.DefaultLogger.Info("Connected to Redis")

	pterm.DefaultLogger.Info("Pinging to Redis")
	_, err = RedisClient.Ping(Ctx).Result()
	if err != nil {
		pterm.DefaultLogger.Error("Unable to ping to Redis")
		os.Exit(1)
	}
	pterm.DefaultLogger.Info("Pinged to Redis")
}
