package main

import (
	"LoadTesterGo/internal/config"
	"LoadTesterGo/internal/logger"
	"LoadTesterGo/internal/master"
)

func main() {
	// Initialize configuration
	cfg := config.LoadConfig()

	// Initialize logger
	logger.InitLogger()

	// Create a new Master
	m := master.NewMaster(cfg.WorkerCount, cfg.TargetURL)

	// Start the distributed load test
	m.StartDistributedLoadTest(cfg.RequestsPerWorker)
}
