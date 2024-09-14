package config

import (
	"log"
	"os"
	"strconv"
)

type Config struct {
	WorkerCount      int
	RequestsPerWorker int
	TargetURL        string
}

func LoadConfig() *Config {
	workerCount, err := strconv.Atoi(os.Getenv("WORKER_COUNT"))
	if err != nil {
		log.Fatal("Invalid WORKER_COUNT")
	}

	requestsPerWorker, err := strconv.Atoi(os.Getenv("REQUESTS_PER_WORKER"))
	if err != nil {
		log.Fatal("Invalid REQUESTS_PER_WORKER")
	}

	targetURL := os.Getenv("TARGET_URL")
	if targetURL == "" {
		log.Fatal("TARGET_URL is required")
	}

	return &Config{
		WorkerCount:      workerCount,
		RequestsPerWorker: requestsPerWorker,
		TargetURL:        targetURL,
	}
}
