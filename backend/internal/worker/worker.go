package worker

import (
	"LoadTesterGo/internal/logger"
	"net/http"
	"time"
)

type Worker struct {
	ID      int
	Target  string
	Results chan<- Result
}

type Result struct {
	WorkerID  int
	Status    int
	Duration  time.Duration
	Error     error
}

func NewWorker(id int, target string, results chan<- Result) *Worker {
	return &Worker{
		ID:      id,
		Target:  target,
		Results: results,
	}
}

func (w *Worker) StartLoadTest(requests int) {
	for i := 0; i < requests; i++ {
		start := time.Now()
		resp, err := http.Get(w.Target)
		duration := time.Since(start)

		result := Result{
			WorkerID: w.ID,
			Status:   http.StatusInternalServerError,
			Duration: duration,
			Error:    err,
		}

		if err == nil {
			result.Status = resp.StatusCode
		} else {
			logger.LogError(w.ID, err)
		}

		w.Results <- result
	}
}
