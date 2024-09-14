package metrics

import (
	"LoadTesterGo/internal/worker"
	"fmt"
	"time"
)

type Metrics struct {
	SuccessCount int
	FailureCount int
	TotalTime    time.Duration
}

func NewMetrics() *Metrics {
	return &Metrics{}
}

func (m *Metrics) ProcessResult(result worker.Result) {
	if result.Status >= 200 && result.Status < 300 {
		m.SuccessCount++
	} else {
		m.FailureCount++
	}
	m.TotalTime += result.Duration
}

func (m *Metrics) Report(workerCount, requestCount int) {
	totalRequests := workerCount * requestCount
	fmt.Printf("Total requests: %d\n", totalRequests)
	fmt.Printf("Success: %d\n", m.SuccessCount)
	fmt.Printf("Failures: %d\n", m.FailureCount)
	avgTime := m.TotalTime / time.Duration(totalRequests)
	fmt.Printf("Average response time: %s\n", avgTime)
}
