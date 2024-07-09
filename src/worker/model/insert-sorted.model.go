package worker_model

import "time"

type InsertSorted struct {
	Data         interface{}    `json:"data"`                    // Data passed to callback.
	ExecuteAt    *time.Time     `json:"execute_at,omitempty"`    // Execution date of callback.
	FutureOffset *time.Duration `json:"future_offset,omitempty"` // Difference between the current time and the execution date in seconds. Will be loaded just if execute_at is not set.
}

// Generates the execution date of the data callback populating ExecuteAt.
//
// If execute_at is not set, it will be set to the current time plus the future_offset in seconds. Otherwise, it will remain the same.
func (i *InsertSorted) GenerateExecutionDate() {
	if i.ExecuteAt != nil {
		return
	}

	currentTime := time.Now()
	executeAt := currentTime.Add(*i.FutureOffset * time.Second)
	i.ExecuteAt = &executeAt
}
