package redis_main

import redis_connect "github.com/Rfluid/http-scheduler/src/redis/connect"

func Redis() {
	redis_connect.Connect()
}
