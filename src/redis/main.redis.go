package redis

import (
	redis_connect "github.com/Rfluid/http-scheduler/src/redis/connect"
	"github.com/redis/go-redis/v9"
)

func Connection() *redis.Client {
	return redis_connect.RedisClient
}
