package env

import (
	"fmt"
	"os"
	"strings"

	"github.com/pterm/pterm"
)

var RedisAddrs []string

func loadRedisEnv() {
	redisAddrs := os.Getenv("REDIS_ADDRS")

	RedisAddrs = strings.Split(redisAddrs, ",")

	pterm.DefaultLogger.Info(
		fmt.Sprintf("Redis environment done with addrs %s", redisAddrs),
	)
}
