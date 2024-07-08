package redis_connect

import (
	"context"
	"os"

	"github.com/Rfluid/http-scheduler/src/config/env"
	"github.com/pterm/pterm"
	"github.com/redis/go-redis/v9"
)

var (
	RedisClient *redis.Client
	ctx         = context.Background() // Context passed to Redis client
)

func Connect() {
	pterm.DefaultLogger.Info("Connecting to redis...")
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     env.RedisAddr,
		Password: env.RedisPassword,
		DB:       env.RedisDB,
	})

	// Test the connection
	_, err := RedisClient.Ping(ctx).Result()
	if err != nil {
		pterm.DefaultLogger.Error("Unable to connect to Redis")
		os.Exit(1)
	}
	pterm.DefaultLogger.Info("Connected to Redis")

	pterm.DefaultLogger.Info("Pinging to Redis")
	_, err = RedisClient.Ping(ctx).Result()
	if err != nil {
		pterm.DefaultLogger.Error("Unable to ping to Redis")
		os.Exit(1)
	}
	pterm.DefaultLogger.Info("Pinged to Redis")
}
