package httpclient

import (
	"fmt"
	"net/http"
	"time"
)

func SendRequest(targetURL string) (*http.Response, error) {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := client.Get(targetURL)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	return resp, nil
}
