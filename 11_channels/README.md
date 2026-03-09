# 11 — Channels

## Concept

Channels are Go's primary mechanism for communication between goroutines. They
are typed, thread-safe conduits for passing data.

```go
ch := make(chan int)     // unbuffered channel
ch := make(chan int, 5)  // buffered channel (capacity 5)

ch <- 42      // send
val := <-ch   // receive
close(ch)     // signal no more values
```

Key behaviors:
- **Unbuffered channels** block the sender until a receiver is ready (and vice
  versa). This provides synchronization.
- **Buffered channels** only block when full (sender) or empty (receiver).
- `range ch` receives values until the channel is closed.
- Always close channels from the **sender** side, never the receiver.

Channel direction can be restricted in function signatures:
- `chan<- int` — send-only
- `<-chan int` — receive-only

This is used extensively in Kubernetes for pipeline patterns, event streaming,
and controller work queues.

## Exercise

Open `exercise.go`. You'll build a data processing pipeline:

1. **`Generate(nums ...int) <-chan int`** — Send all numbers into a channel and
   close it. Returns a receive-only channel.

2. **`Square(in <-chan int) <-chan int`** — Read numbers from `in`, square each
   one, and send to an output channel. Close the output when input is exhausted.

3. **`FilterChan(in <-chan int, keep func(int) bool) <-chan int`** — Read numbers
   from `in`, only forward those where `keep` returns true.

4. **`Merge(channels ...<-chan int) <-chan int`** — Fan-in: merge multiple input
   channels into a single output channel. All inputs are read concurrently.

5. Wire them together: Generate → Square → Filter (keep even) → Print.
   Also demonstrate Merge by combining two generators.

Run with: `go run exercise.go`

## Hint

- Each pipeline stage runs in its own goroutine. The function launches the
  goroutine and immediately returns the output channel.
- Always `close()` the output channel when a stage finishes, or the downstream
  `range` will deadlock.
- For `Merge`, launch one goroutine per input channel, use a WaitGroup, and
  close the output channel when all inputs are drained.
- Channel direction types (`<-chan`, `chan<-`) prevent accidental misuse.
