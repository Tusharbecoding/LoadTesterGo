package main

import (
	"LoadTesterGo/internal/config"
	"LoadTesterGo/internal/master"
	"log"

	"github.com/joho/godotenv"
)

func main() {
    // Load .env file
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    // Load config and start the load test
    cfg := config.LoadConfig()

    // Initialize Master and start distributed load test
    m := master.NewMaster(cfg.WorkerCount, cfg.TargetURL)
    m.StartDistributedLoadTest(cfg.RequestsPerWorker)
}
