# 09 — Goroutines and WaitGroups

## Concept

Goroutines are lightweight concurrent functions launched with the `go` keyword.
They are much cheaper than OS threads — you can run thousands simultaneously.

```go
go myFunction()       // runs myFunction concurrently
go func() { ... }()  // anonymous goroutine
```

The main goroutine won't wait for others to finish. Use `sync.WaitGroup` to
coordinate:

```go
var wg sync.WaitGroup
wg.Add(1)            // tell WaitGroup to expect one goroutine
go func() {
    defer wg.Done()  // signal completion when done
    doWork()
}()
wg.Wait()            // block until all goroutines call Done()
```

`defer` schedules a call to run when the enclosing function returns — perfect
for cleanup and `wg.Done()`.

In Kubernetes, goroutines power controllers, informers, and the API server's
request handling. Understanding them is essential.

## Exercise

Open `exercise.go`. You'll simulate a parallel API fetcher:

1. **Define a `FetchResult` struct** with `URL string`, `Status string`, and
   `Duration time.Duration`.

2. **`FetchURL(url string)`** — Simulate fetching a URL with a random delay
   (100-500ms). Returns a `FetchResult`. The Status should be "200 OK" for
   URLs containing "good" and "500 Error" for URLs containing "bad", and
   "404 Not Found" otherwise.

3. **`FetchAll(urls []string)`** — Fetch all URLs concurrently using goroutines
   and a WaitGroup. Collect results into a slice. Return the slice once all
   goroutines complete. Results must be collected safely (think about shared
   state).

4. **`FetchSequential(urls []string)`** — Fetch all URLs one at a time (for
   comparison). The main function will compare timings.

Run with: `go run exercise.go`

## Hint

- Use `time.Sleep(time.Duration(rand.Intn(400)+100) * time.Millisecond)` for
  random delays.
- Use `time.Now()` and `time.Since(start)` to measure duration.
- To safely collect results from goroutines into a shared slice, you'll need
  either a mutex or a pre-allocated slice with index-based assignment.
- Pre-allocating: `results := make([]FetchResult, len(urls))` and having
  goroutine `i` write to `results[i]` is safe without a mutex because each
  goroutine writes to a different index.
