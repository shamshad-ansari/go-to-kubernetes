package main

import (
	"fmt"
	"math/rand"
	"time"
)

// HealthResult holds the outcome of a single health check.
type HealthResult struct {
	ServiceName string
	Healthy     bool
	Message     string
	Latency     time.Duration
}

// HealthChecker is the interface that all checkable services must implement.
type HealthChecker interface {
	Name() string
	CheckHealth() HealthResult
}

// --- Service Implementations ---

// APIServer simulates a Kubernetes API server health check.
// 80% chance of being healthy, 200-500ms latency.
type APIServer struct{}

func (a APIServer) Name() string { return "api-server" }

// CheckHealth simulates an API server health check.
// 80% healthy, 200-500ms latency.
// Healthy: "all endpoints responding", Unhealthy: "high latency on /apis"
func (a APIServer) CheckHealth() HealthResult {
	// TODO: Record start time with time.Now().
	// TODO: Sleep for rand.Intn(300)+200 milliseconds.
	// TODO: Determine healthy with rand.Float64() < 0.8.
	// TODO: Set message based on health status.
	// TODO: Return HealthResult with ServiceName from Name(), Healthy,
	//       Message, and Latency from time.Since(start).

	return HealthResult{} // replace this
}

// Database simulates a database health check.
// 70% chance of being healthy, 300-800ms latency.
type Database struct {
	ConnectionString string
}

func (d Database) Name() string { return "database" }

// CheckHealth simulates a database health check.
// 70% healthy, 300-800ms latency.
// Healthy: "connections available", Unhealthy: "connection pool exhausted"
func (d Database) CheckHealth() HealthResult {
	// TODO: Same pattern as APIServer but with different parameters.
	// Sleep: rand.Intn(500)+300, healthy: rand.Float64() < 0.7

	return HealthResult{} // replace this
}

// Cache simulates a cache (like Redis) health check.
// 90% chance of being healthy, 50-200ms latency.
type Cache struct{}

func (c Cache) Name() string { return "cache" }

// CheckHealth simulates a cache health check.
// 90% healthy, 50-200ms latency.
// Healthy: "hit rate normal", Unhealthy: "eviction rate too high"
func (c Cache) CheckHealth() HealthResult {
	// TODO: Same pattern. Sleep: rand.Intn(150)+50, healthy: rand.Float64() < 0.9

	return HealthResult{} // replace this
}

// RunHealthChecks probes all services concurrently and collects results.
//
// Strategy:
//  1. Create a HealthResult channel
//  2. Launch one goroutine per service that calls CheckHealth() and sends
//     the result into the channel
//  3. Receive exactly len(services) results from the channel
//  4. Return the collected results
func RunHealthChecks(services []HealthChecker) []HealthResult {
	// TODO: Create a channel of HealthResult.
	// TODO: Launch a goroutine per service.
	// TODO: Collect all results from the channel.
	// TODO: Return the results slice.

	return nil // replace this
}

// PrintReport prints a formatted health report table.
//
// Example output:
//
//	SERVICE         STATUS    LATENCY     MESSAGE
//	api-server      HEALTHY   234ms       all endpoints responding
//	database        UNHEALTHY 567ms       connection pool exhausted
//	cache           HEALTHY   89ms        hit rate normal
func PrintReport(results []HealthResult) {
	// TODO: Print a header line.
	// TODO: Loop through results and print each one.
	//       Use "HEALTHY" or "UNHEALTHY" based on the Healthy field.
	//       Align columns using fmt.Printf width specifiers.
	//       Example: fmt.Printf("  %-15s %-10s %-12s %s\n", ...)
}

// Ensure imports are used.
var (
	_ = rand.Float64
	_ = time.Now
)

func main() {
	services := []HealthChecker{
		APIServer{},
		Database{ConnectionString: "postgres://localhost:5432/mydb"},
		Cache{},
	}

	fmt.Println("=== Kubernetes Health Probe ===")
	fmt.Println("Checking services concurrently...\n")

	start := time.Now()
	results := RunHealthChecks(services)
	elapsed := time.Since(start)

	PrintReport(results)

	healthy := 0
	for _, r := range results {
		if r.Healthy {
			healthy++
		}
	}

	fmt.Printf("\n  Summary: %d/%d healthy | Total probe time: %v\n",
		healthy, len(results), elapsed)

	if elapsed < 900*time.Millisecond {
		fmt.Println("  (Probes ran concurrently — total time < slowest individual probe)")
	}
}
