package worker_service

import (
	"sync"

	worker_model "github.com/Rfluid/http-scheduler/src/worker/model"
)

func InsertManySorted(body []worker_model.InsertSorted) error {
	var wg sync.WaitGroup
	errCh := make(chan error, len(body))
	wg.Add(len(body))

	go func() {
		for _, b := range body {
			go func(b worker_model.InsertSorted) {
				err := InsertSorted(b)
				errCh <- err
			}(b)
		}
	}()

	go func() {
		defer close(errCh)
		wg.Wait()
	}()

	for err := range errCh {
		wg.Done()
		if err != nil {
			return err
		}
	}

	return nil
}
