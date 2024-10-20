package generator

import (
	"ddos/internal/client"
	"ddos/internal/models"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

type Generator struct {
	HTTPClient *client.Client
}

func (g *Generator) Attack(req models.LoadTestRequest) (string, error) {
	var wg sync.WaitGroup
	results := make(chan string, req.NumRequests)

	for i := 0; i < req.NumRequests; i++ {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()
			result := sendRequest(req.URL)
			results <- result
		}(i)

		time.Sleep(time.Duration(req.Delay) * time.Millisecond)
	}

	wg.Wait()
	close(results)

	aggregatedResults := "Results:\n"
	for result := range results {
		aggregatedResults += result + "\n"
	}

	return aggregatedResults, nil
}

func sendRequest(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Error sending request to %s: %v", url, err)
		return fmt.Sprintf("Failed to send request to %s", url)
	}
	defer resp.Body.Close()

	return fmt.Sprintf("Request to %s returned status: %s", url, resp.Status)
}

// sendRequest sends an HTTP GET request to the specified URL and returns the result
func (g *Generator) sendRequest(url string) string {
	resp, err := g.HTTPClient.SendRequest(url)
	if err != nil {
		log.Printf("Error sending request to %s: %v", url, err)
		return fmt.Sprintf("Failed to send request to %s", url)
	}
	defer resp.Body.Close()

	return fmt.Sprintf("Request to %s returned status: %s", url, resp.Status)
}
