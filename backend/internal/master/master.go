package master

import (
	"LoadTesterGo/internal/metrics"
	"LoadTesterGo/internal/worker"
	"sync"
)

type Master struct {
	Workers          []*worker.Worker
	Results          chan worker.Result
	TargetURL        string
	metricsCollector *metrics.Metrics
}

func NewMaster(workerCount int, targetURL string) *Master {
	results := make(chan worker.Result, workerCount*100)
	m := &Master{
		Workers:          make([]*worker.Worker, workerCount),
		Results:          results,
		TargetURL:        targetURL,
		metricsCollector: metrics.NewMetrics(),
	}

	for i := 0; i < workerCount; i++ {
		m.Workers[i] = worker.NewWorker(i, targetURL, results)
	}
	return m
}

func (m *Master) StartDistributedLoadTest(requestsPerWorker int) {
	var wg sync.WaitGroup

	for _, w := range m.Workers {
		wg.Add(1)
		go func(w *worker.Worker) {
			defer wg.Done()
			w.StartLoadTest(requestsPerWorker)
		}(w)
	}

	wg.Wait()
	close(m.Results)

	for result := range m.Results {
		m.metricsCollector.ProcessResult(result)
	}
}

func (m *Master) GetMetrics() *metrics.Metrics {
	return m.metricsCollector
}
