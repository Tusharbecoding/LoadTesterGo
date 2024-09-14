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
	TotalRequests int
}

func NewMetrics() *Metrics {
	return &Metrics{}
}

// ProcessResult processes each result from the workers
func (m *Metrics) ProcessResult(result worker.Result) {
	if result.Status >= 200 && result.Status < 300 {
		m.SuccessCount++
	} else {
		m.FailureCount++
	}
	m.TotalTime += result.Duration
	m.TotalRequests++
}

// AvgResponseTime calculates and returns the average response time for all requests
func (m *Metrics) AvgResponseTime() time.Duration {
	if m.TotalRequests == 0 {
		return 0
	}
	return m.TotalTime / time.Duration(m.TotalRequests)
}

// Report prints the final metrics report to the console
func (m *Metrics) Report(workerCount, requestCount int) {
	totalRequests := workerCount * requestCount
	fmt.Printf("Total requests: %d\n", totalRequests)
	fmt.Printf("Success: %d\n", m.SuccessCount)
	fmt.Printf("Failures: %d\n", m.FailureCount)
	avgTime := m.AvgResponseTime()
	fmt.Printf("Average response time: %s\n", avgTime)
}
