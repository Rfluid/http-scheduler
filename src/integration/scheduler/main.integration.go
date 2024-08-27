package scheduler

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

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
	Worker.SetCallback(workerCallback)
	Worker.SetTryDequeueErrCallback(tryDequeueErrCallback)
	Worker.TryDequeue(local_redis.Context()) // Reschedule after restart
	pterm.DefaultLogger.Info("Scheduler integration loaded")
}

func workerCallback(data redis.Z) error {
	if env.PropagateUrl == "" {
		pterm.DefaultLogger.Info(fmt.Sprintf("Propagating %v", data.Member))
		return nil
	}

	jsonData, err := json.Marshal(data.Member)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(env.PropagateMethod, env.PropagateUrl, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	req.Header = http.Header{env.PropagateHeaderKey: {env.PropagateHeaderValue}}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		return errors.New(string(bodyBytes))
	}

	return nil
}

func tryDequeueErrCallback(err error) {
	if err != nil {
		pterm.DefaultLogger.Error(err.Error())
	}
}
