//go:build ignore

package main

import (
	"fmt"
	"math/rand"
	"time"
)

type HealthResult struct {
	ServiceName string
	Healthy     bool
	Message     string
	Latency     time.Duration
}

type HealthChecker interface {
	Name() string
	CheckHealth() HealthResult
}

type APIServer struct{}

func (a APIServer) Name() string { return "api-server" }

func (a APIServer) CheckHealth() HealthResult {
	start := time.Now()
	time.Sleep(time.Duration(rand.Intn(300)+200) * time.Millisecond)

	healthy := rand.Float64() < 0.8
	msg := "all endpoints responding"
	if !healthy {
		msg = "high latency on /apis"
	}

	return HealthResult{
		ServiceName: a.Name(),
		Healthy:     healthy,
		Message:     msg,
		Latency:     time.Since(start),
	}
}

type Database struct {
	ConnectionString string
}

func (d Database) Name() string { return "database" }

func (d Database) CheckHealth() HealthResult {
	start := time.Now()
	time.Sleep(time.Duration(rand.Intn(500)+300) * time.Millisecond)

	healthy := rand.Float64() < 0.7
	msg := "connections available"
	if !healthy {
		msg = "connection pool exhausted"
	}

	return HealthResult{
		ServiceName: d.Name(),
		Healthy:     healthy,
		Message:     msg,
		Latency:     time.Since(start),
	}
}

type Cache struct{}

func (c Cache) Name() string { return "cache" }

func (c Cache) CheckHealth() HealthResult {
	start := time.Now()
	time.Sleep(time.Duration(rand.Intn(150)+50) * time.Millisecond)

	healthy := rand.Float64() < 0.9
	msg := "hit rate normal"
	if !healthy {
		msg = "eviction rate too high"
	}

	return HealthResult{
		ServiceName: c.Name(),
		Healthy:     healthy,
		Message:     msg,
		Latency:     time.Since(start),
	}
}

func RunHealthChecks(services []HealthChecker) []HealthResult {
	ch := make(chan HealthResult, len(services))

	for _, svc := range services {
		go func(s HealthChecker) {
			ch <- s.CheckHealth()
		}(svc)
	}

	results := make([]HealthResult, 0, len(services))
	for i := 0; i < len(services); i++ {
		results = append(results, <-ch)
	}
	return results
}

func PrintReport(results []HealthResult) {
	fmt.Printf("  %-15s %-10s %-12s %s\n", "SERVICE", "STATUS", "LATENCY", "MESSAGE")
	fmt.Printf("  %-15s %-10s %-12s %s\n", "-------", "------", "-------", "-------")

	for _, r := range results {
		status := "HEALTHY"
		if !r.Healthy {
			status = "UNHEALTHY"
		}
		fmt.Printf("  %-15s %-10s %-12v %s\n",
			r.ServiceName, status, r.Latency.Round(time.Millisecond), r.Message)
	}
}

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
