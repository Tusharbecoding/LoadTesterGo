package main

import (
	"LoadTesterGo/internal/config"
	"LoadTesterGo/internal/master"
	"encoding/json"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

type LoadTestRequest struct {
	TargetURL string `json:"target_url"`
}

type LoadTestResponse struct {
	TotalRequests     int    `json:"totalRequests"`
	SuccessCount      int    `json:"successCount"`
	FailureCount      int    `json:"failureCount"`
	AvgResponseTime   string `json:"avgResponseTime"`
}

func loadTestHandler(w http.ResponseWriter, r *http.Request) {
	var requestData LoadTestRequest

	// Parse the incoming JSON payload
	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil || requestData.TargetURL == "" {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Initialize the load tester with the provided URL
	cfg := &config.Config{
		WorkerCount:      5,
		RequestsPerWorker: 100,
		TargetURL:        requestData.TargetURL,
	}

	m := master.NewMaster(cfg.WorkerCount, cfg.TargetURL)
	m.StartDistributedLoadTest(cfg.RequestsPerWorker)

	// Create the response with load testing metrics
	response := LoadTestResponse{
		TotalRequests:   cfg.WorkerCount * cfg.RequestsPerWorker,
		SuccessCount:    m.GetMetrics().SuccessCount,
		FailureCount:    m.GetMetrics().FailureCount,
		AvgResponseTime: m.GetMetrics().AvgResponseTime().String(),
	}

	// Send the response back to the frontend
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}


func main() {
	// Load environment variables from .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Set up the API endpoint
	http.HandleFunc("/load-test", loadTestHandler)

	// Start the server
	log.Println("Server running on port 8080...")
	http.ListenAndServe(":8080", nil)
}
