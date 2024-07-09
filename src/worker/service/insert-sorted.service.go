package worker_service

import (
	"github.com/Rfluid/http-scheduler/src/integration/scheduler"
	"github.com/Rfluid/http-scheduler/src/redis"
	worker_model "github.com/Rfluid/http-scheduler/src/worker/model"
)

func InsertSorted(body worker_model.InsertSorted) error {
	body.GenerateExecutionDate()
	err := scheduler.Worker.InsertSortedByDate(body.Data, *body.ExecuteAt, redis.Context())

	return err
}
