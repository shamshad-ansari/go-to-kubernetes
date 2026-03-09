package main

import (
	"fmt"
	"math/rand"
	"strings"
	"sync"
	"time"
)

// FetchResult holds the outcome of a simulated URL fetch.
type FetchResult struct {
	URL      string
	Status   string
	Duration time.Duration
}

// FetchURL simulates an HTTP fetch with a random delay (100-500ms).
//
// Status logic:
//   - URL contains "good" → "200 OK"
//   - URL contains "bad"  → "500 Error"
//   - Otherwise           → "404 Not Found"
func FetchURL(url string) FetchResult {
	// TODO: Record the start time.
	// TODO: Sleep for a random duration between 100-500ms.
	// TODO: Determine the status based on the URL contents.
	// TODO: Return a FetchResult with URL, Status, and Duration.

	return FetchResult{} // replace this
}

// FetchAll fetches all URLs concurrently using goroutines and a WaitGroup.
// It returns all results once every fetch is complete.
//
// Safety: since each goroutine writes to a different index in a pre-allocated
// slice, no mutex is needed.
func FetchAll(urls []string) []FetchResult {
	// TODO: Pre-allocate a results slice: make([]FetchResult, len(urls))
	// TODO: Create a sync.WaitGroup.
	// TODO: For each URL, Add(1) to the WaitGroup and launch a goroutine
	//       that calls FetchURL and stores the result at the correct index.
	//       Don't forget defer wg.Done() inside the goroutine.
	// TODO: Wait for all goroutines to finish.
	// TODO: Return results.

	return nil // replace this
}

// FetchSequential fetches all URLs one at a time (no concurrency).
func FetchSequential(urls []string) []FetchResult {
	// TODO: Create a results slice.
	// TODO: Loop through URLs, calling FetchURL for each, and append results.
	// TODO: Return results.

	return nil // replace this
}

// Ensure imports are used.
var (
	_ = rand.Intn
	_ = strings.Contains
	_ = sync.WaitGroup{}
)

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
