package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Result holds the outcome of a job execution.
type Result struct {
	JobID    string
	Output   string
	Err      error
	Duration time.Duration
}

// Job defines the interface for executable work items.
type Job interface {
	ID() string
	Execute() Result
}

// WorkerPool manages a fixed number of workers processing jobs from a queue.
type WorkerPool struct {
	numWorkers int
	jobs       chan Job
	results    chan Result
	wg         sync.WaitGroup
}

// NewWorkerPool creates a new worker pool.
//   - numWorkers: how many concurrent workers to run
//   - jobBuffer:  buffer size for the jobs channel
func NewWorkerPool(numWorkers, jobBuffer int) *WorkerPool {
	// TODO: Create and return a WorkerPool with initialized channels.
	// The results channel should be buffered too (same size as jobBuffer).

	return nil // replace this
}

// Start launches all worker goroutines. Each worker:
//  1. Ranges over the jobs channel
//  2. Calls job.Execute() for each job
//  3. Sends the result to the results channel
//  4. Calls wg.Done() when the jobs channel is closed and drained
func (wp *WorkerPool) Start() {
	// TODO: For each worker (up to numWorkers):
	//   - wg.Add(1)
	//   - Launch a goroutine that:
	//     - Defers wg.Done()
	//     - Ranges over wp.jobs
	//     - Calls Execute() and sends result to wp.results
}

// Submit sends a job to the pool for processing.
func (wp *WorkerPool) Submit(job Job) {
	// TODO: Send the job to the jobs channel.
}

// Results returns the results channel (receive-only).
func (wp *WorkerPool) Results() <-chan Result {
	// TODO: Return wp.results.

	return nil // replace this
}

// Shutdown gracefully shuts down the pool:
//  1. Close the jobs channel (no more jobs accepted)
//  2. Wait for all workers to finish processing
//  3. Close the results channel
func (wp *WorkerPool) Shutdown() {
	// TODO: close(wp.jobs)
	// TODO: wp.wg.Wait()
	// TODO: close(wp.results)
}

// --- Job Implementations ---

// ImagePullJob simulates pulling a container image.
// Takes 200-600ms. Fails 10% of the time.
type ImagePullJob struct {
	Image string
}

func (j ImagePullJob) ID() string { return "pull:" + j.Image }

// Execute simulates pulling a container image.
// 200-600ms latency, 10% failure rate.
// Success: "pulled <image> successfully"
// Failure error: "failed to pull image: <image>"
func (j ImagePullJob) Execute() Result {
	// TODO: Record start time.
	// TODO: Sleep rand.Intn(400)+200 ms.
	// TODO: 10% chance of error (rand.Float64() < 0.1).
	// TODO: Return Result with JobID from j.ID(), Output or Err, and Duration.

	return Result{} // replace this
}

// HealthCheckJob simulates running a health check on a service.
// Takes 100-400ms. Fails 20% of the time.
type HealthCheckJob struct {
	Service string
}

func (j HealthCheckJob) ID() string { return "health:" + j.Service }

// Execute simulates a health check on a service.
// 100-400ms latency, 20% failure rate.
// Success: "<service> is healthy"
// Failure error: "<service> is unhealthy"
func (j HealthCheckJob) Execute() Result {
	// TODO: Same pattern as ImagePullJob with different parameters.
	// Sleep: rand.Intn(300)+100, failure: rand.Float64() < 0.2

	return Result{} // replace this
}

// Ensure imports are used.
var (
	_ = rand.Intn
	_ = sync.WaitGroup{}
	_ = time.Now
)

func main() {
	pool := NewWorkerPool(3, 10)
	pool.Start()

	jobs := []Job{
		ImagePullJob{Image: "nginx:latest"},
		ImagePullJob{Image: "redis:7"},
		ImagePullJob{Image: "postgres:15"},
		HealthCheckJob{Service: "api-server"},
		HealthCheckJob{Service: "scheduler"},
		HealthCheckJob{Service: "etcd"},
		ImagePullJob{Image: "golang:1.21"},
		HealthCheckJob{Service: "controller-manager"},
	}

	fmt.Printf("=== Worker Pool (3 workers, %d jobs) ===\n\n", len(jobs))

	for _, j := range jobs {
		pool.Submit(j)
	}

	go func() {
		pool.Shutdown()
	}()

	succeeded, failed := 0, 0
	for r := range pool.Results() {
		if r.Err != nil {
			fmt.Printf("  FAIL  %-30s  %v  (%v)\n", r.JobID, r.Err, r.Duration.Round(time.Millisecond))
			failed++
		} else {
			fmt.Printf("  OK    %-30s  %s  (%v)\n", r.JobID, r.Output, r.Duration.Round(time.Millisecond))
			succeeded++
		}
	}

	fmt.Printf("\n=== Results: %d succeeded, %d failed, %d total ===\n",
		succeeded, failed, succeeded+failed)
}
