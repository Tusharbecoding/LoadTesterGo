package master

import (
	"LoadTesterGo/internal/metrics"
	"LoadTesterGo/internal/worker"
	"sync"
)

type Master struct {
	Workers   []*worker.Worker
	Results   chan worker.Result
	TargetURL string
}

func NewMaster(workerCount int, targetURL string) *Master {
	results := make(chan worker.Result, workerCount*100)
	m := &Master{
		Workers:   make([]*worker.Worker, workerCount),
		Results:   results,
		TargetURL: targetURL,
	}

	for i := 0; i < workerCount; i++ {
		m.Workers[i] = worker.NewWorker(i, targetURL, results)
	}
	return m
}

func (m *Master) StartDistributedLoadTest(requestsPerWorker int) {
	var wg sync.WaitGroup

	// Start workers to simulate load
	for _, w := range m.Workers {
		wg.Add(1)
		go func(w *worker.Worker) {
			defer wg.Done()
			w.StartLoadTest(requestsPerWorker)
		}(w)
	}

	// Wait for all workers to finish
	wg.Wait()
	close(m.Results)

	// Process metrics
	m.ProcessMetrics(len(m.Workers), requestsPerWorker)
}

func (m *Master) ProcessMetrics(workerCount, requestsPerWorker int) {
	metricsCollector := metrics.NewMetrics()

	for result := range m.Results {
		metricsCollector.ProcessResult(result)
	}

	metricsCollector.Report(workerCount, requestsPerWorker)
}
