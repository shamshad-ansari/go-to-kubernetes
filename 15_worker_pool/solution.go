//go:build ignore

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Result struct {
	JobID    string
	Output   string
	Err      error
	Duration time.Duration
}

type Job interface {
	ID() string
	Execute() Result
}

type WorkerPool struct {
	numWorkers int
	jobs       chan Job
	results    chan Result
	wg         sync.WaitGroup
}

func NewWorkerPool(numWorkers, jobBuffer int) *WorkerPool {
	return &WorkerPool{
		numWorkers: numWorkers,
		jobs:       make(chan Job, jobBuffer),
		results:    make(chan Result, jobBuffer),
	}
}

func (wp *WorkerPool) Start() {
	for i := 0; i < wp.numWorkers; i++ {
		wp.wg.Add(1)
		go func() {
			defer wp.wg.Done()
			for job := range wp.jobs {
				wp.results <- job.Execute()
			}
		}()
	}
}

func (wp *WorkerPool) Submit(job Job) {
	wp.jobs <- job
}

func (wp *WorkerPool) Results() <-chan Result {
	return wp.results
}

func (wp *WorkerPool) Shutdown() {
	close(wp.jobs)
	wp.wg.Wait()
	close(wp.results)
}

type ImagePullJob struct {
	Image string
}

func (j ImagePullJob) ID() string { return "pull:" + j.Image }

func (j ImagePullJob) Execute() Result {
	start := time.Now()
	time.Sleep(time.Duration(rand.Intn(400)+200) * time.Millisecond)

	if rand.Float64() < 0.1 {
		return Result{
			JobID:    j.ID(),
			Err:      fmt.Errorf("failed to pull image: %s", j.Image),
			Duration: time.Since(start),
		}
	}

	return Result{
		JobID:    j.ID(),
		Output:   fmt.Sprintf("pulled %s successfully", j.Image),
		Duration: time.Since(start),
	}
}

type HealthCheckJob struct {
	Service string
}

func (j HealthCheckJob) ID() string { return "health:" + j.Service }

func (j HealthCheckJob) Execute() Result {
	start := time.Now()
	time.Sleep(time.Duration(rand.Intn(300)+100) * time.Millisecond)

	if rand.Float64() < 0.2 {
		return Result{
			JobID:    j.ID(),
			Err:      fmt.Errorf("%s is unhealthy", j.Service),
			Duration: time.Since(start),
		}
	}

	return Result{
		JobID:    j.ID(),
		Output:   fmt.Sprintf("%s is healthy", j.Service),
		Duration: time.Since(start),
	}
}

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
