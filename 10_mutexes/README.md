# 10 — Mutexes

## Concept

When multiple goroutines read and write shared data, you get **race conditions**
— unpredictable bugs where the outcome depends on timing. Mutexes prevent this
by ensuring only one goroutine accesses the critical section at a time.

Go provides two mutex types:

- **`sync.Mutex`** — Exclusive lock. One goroutine at a time, period.
- **`sync.RWMutex`** — Allows multiple simultaneous readers, but only one writer.
  Better performance when reads vastly outnumber writes.

```go
var mu sync.Mutex

mu.Lock()
// critical section — only one goroutine here at a time
mu.Unlock()
```

```go
var rw sync.RWMutex

rw.RLock()    // multiple goroutines can hold this simultaneously
// read-only section
rw.RUnlock()

rw.Lock()     // exclusive — blocks all readers and writers
// write section
rw.Unlock()
```

In Kubernetes, shared caches, informer stores, and controller queues all use
mutexes. Getting locking right is critical — too broad kills performance, too
narrow causes races.

## Exercise

Open `exercise.go`. You'll build a thread-safe scoreboard:

1. **Define a `Scoreboard` struct** with a `scores map[string]int` and a
   `sync.RWMutex`.

2. **`NewScoreboard()`** — Constructor that initializes the map.

3. **`RecordScore(name string, points int)`** — Add points to a player's score.
   Must be write-locked.

4. **`GetScore(name string) (int, bool)`** — Get a player's score. Must be
   read-locked. Returns the score and whether the player exists.

5. **`TopScorers(n int) []string`** — Return the top N player names sorted by
   score (descending). Must be read-locked.

6. **Test it** with 100 concurrent goroutines recording scores and verify the
   final totals are correct.

Run with: `go run exercise.go`

## Hint

- Always `defer mu.Unlock()` right after `mu.Lock()` to prevent deadlocks.
- Use `RLock()`/`RUnlock()` for read-only operations (`GetScore`, `TopScorers`).
- Use `Lock()`/`Unlock()` for write operations (`RecordScore`).
- For sorting, use `sort.Slice` with a custom comparison function.
- To verify correctness: if 100 goroutines each add 1 point to "alice", her
  final score should be exactly 100.
