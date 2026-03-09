//go:build ignore

package main

import (
	"fmt"
	"math/rand"
	"strings"
	"sync"
	"time"
)

type FetchResult struct {
	URL      string
	Status   string
	Duration time.Duration
}

func FetchURL(url string) FetchResult {
	start := time.Now()
	time.Sleep(time.Duration(rand.Intn(400)+100) * time.Millisecond)

	var status string
	switch {
	case strings.Contains(url, "good"):
		status = "200 OK"
	case strings.Contains(url, "bad"):
		status = "500 Error"
	default:
		status = "404 Not Found"
	}

	return FetchResult{
		URL:      url,
		Status:   status,
		Duration: time.Since(start),
	}
}

func FetchAll(urls []string) []FetchResult {
	results := make([]FetchResult, len(urls))
	var wg sync.WaitGroup

	for i, url := range urls {
		wg.Add(1)
		go func(idx int, u string) {
			defer wg.Done()
			results[idx] = FetchURL(u)
		}(i, url)
	}

	wg.Wait()
	return results
}

func FetchSequential(urls []string) []FetchResult {
	results := make([]FetchResult, 0, len(urls))
	for _, url := range urls {
		results = append(results, FetchURL(url))
	}
	return results
}

func main() {
	urls := []string{
		"https://api.good-service.com/health",
		"https://api.bad-service.com/health",
		"https://api.good-service.com/data",
		"https://api.unknown.com/page",
		"https://api.good-service.com/users",
		"https://api.bad-service.com/crash",
	}

	fmt.Println("=== Sequential Fetching ===")
	start := time.Now()
	seqResults := FetchSequential(urls)
	seqDuration := time.Since(start)
	for _, r := range seqResults {
		fmt.Printf("  %-45s → %-15s (%v)\n", r.URL, r.Status, r.Duration)
	}
	fmt.Printf("  Total time: %v\n", seqDuration)

	fmt.Println("\n=== Concurrent Fetching ===")
	start = time.Now()
	conResults := FetchAll(urls)
	conDuration := time.Since(start)
	for _, r := range conResults {
		fmt.Printf("  %-45s → %-15s (%v)\n", r.URL, r.Status, r.Duration)
	}
	fmt.Printf("  Total time: %v\n", conDuration)

	fmt.Printf("\n=== Speedup: %.1fx faster with concurrency ===\n",
		float64(seqDuration)/float64(conDuration))
}
