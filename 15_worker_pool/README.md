# 15 — Capstone: Worker Pool (Kubernetes Job Scheduler)

## Context

In Kubernetes, controllers process work items from a queue. The **workqueue**
package provides rate-limited, deduplicated queues. Controllers launch a fixed
number of worker goroutines that pop items from the queue, process them, and
report results.

This is the **worker pool** pattern: a fixed number of workers pulling jobs from
a shared channel. It's one of the most important concurrency patterns in Go
and is used throughout Kubernetes (controller workers, admission webhooks, garbage
collectors, etc.).

## Concepts Combined

- **Interfaces** — Job defines the contract for executable work
- **Structs** — Job implementations and the pool itself
- **Goroutines** — Workers run as goroutines
- **Channels** — Job queue and result collection
- **WaitGroups** — Coordinating worker shutdown
- **Generics** — Optional: make the pool generic over result types
- **Error handling** — Jobs can succeed or fail

## Exercise

Open `exercise.go`. You'll build a complete worker pool system:

1. **Define a `Job` interface** with:
   - `ID() string`
   - `Execute() Result`

2. **Define a `Result` struct** with `JobID string`, `Output string`,
   `Err error`, `Duration time.Duration`.

3. **Define a `WorkerPool` struct** with:
   - `numWorkers int`
   - `jobs       chan Job`
   - `results    chan Result`
   - `wg         sync.WaitGroup`

4. **`NewWorkerPool(numWorkers, jobBuffer int)`** — Constructor.

5. **`Start()`** — Launch numWorkers goroutines, each reading from the jobs
   channel and writing to results.

6. **`Submit(job Job)`** — Send a job to the pool.

7. **`Results() <-chan Result`** — Return the results channel (receive-only).

8. **`Shutdown()`** — Close the jobs channel and wait for all workers to finish,
   then close the results channel.

9. **Implement two job types:**
   - `ImagePullJob` — simulates pulling a container image (200-600ms)
   - `HealthCheckJob` — simulates checking a service (100-400ms)

Run with: `go run exercise.go`

## Hint

- Each worker is a goroutine that `range`s over the `jobs` channel. When the
  channel is closed, the range loop exits.
- `Submit` sends to the jobs channel. Don't close it in Submit.
- `Shutdown` closes `jobs`, waits (`wg.Wait()`), then closes `results`.
- The caller reads results by ranging over `pool.Results()`.
- In the worker, always `defer wg.Done()`.
- This is the exact same pattern used in Kubernetes controllers — study it well.
