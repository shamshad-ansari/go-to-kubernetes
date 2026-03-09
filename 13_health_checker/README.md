# 13 — Capstone: Kubernetes Health Checker

## Context

In Kubernetes, every pod can define **liveness** and **readiness** probes. The
kubelet periodically checks whether containers are alive and ready to serve
traffic. If a liveness probe fails, the container is restarted. If a readiness
probe fails, the pod is removed from the service's endpoints.

This exercise simulates that pattern: you'll build a concurrent health checker
that probes multiple services in parallel and produces a report.

## Concepts Combined

- **Interfaces** — `HealthChecker` defines the contract for checkable services
- **Structs and methods** — Each service type implements the interface
- **Goroutines** — Services are checked concurrently
- **Channels** — Results are collected via a channel
- **Error handling** — Health checks can fail with descriptive errors

## Exercise

Open `exercise.go`. You'll build a health probe system:

1. **Define a `HealthResult` struct** with `ServiceName string`, `Healthy bool`,
   `Message string`, and `Latency time.Duration`.

2. **Define a `HealthChecker` interface** with:
   - `Name() string`
   - `CheckHealth() HealthResult`

3. **Implement three services:**
   - `APIServer` — 80% chance healthy, 200-500ms latency
   - `Database` — 70% chance healthy, 300-800ms latency
   - `Cache` — 90% chance healthy, 50-200ms latency

4. **`RunHealthChecks(services []HealthChecker) []HealthResult`** — Check all
   services concurrently using goroutines and collect results through a channel.

5. **`PrintReport(results []HealthResult)`** — Print a formatted health report
   table.

Run with: `go run exercise.go`

## Hint

- Each service's `CheckHealth()` should sleep for a random duration to simulate
  latency, then use `rand.Float64()` to decide if it's healthy.
- For `RunHealthChecks`, create a channel of `HealthResult`. Launch a goroutine
  per service. Receive exactly `len(services)` results from the channel.
- No WaitGroup needed here — you know exactly how many results to expect.
- Format output with `fmt.Printf` using fixed-width fields for alignment.
