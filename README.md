# Go to Kubernetes — Hands-On Exercise Track

A structured set of Go exercises designed for experienced programmers learning Go
for the first time, with the goal of contributing to open-source Go projects like
[Kubernetes](https://github.com/kubernetes/kubernetes).

## How This Repo Works

Each numbered folder contains one exercise focused on a single Go concept. They
are meant to be completed **in order** — later exercises build on earlier ones.

```
01_variables/       → Variables, types, zero values, conversions
02_functions/       → Multiple returns, error handling patterns
03_slices/          → Slice operations: filter, deduplicate, rotate
04_maps/            → Map lookups, merging, inverting
05_strings/         → Runes, Unicode-safe operations, strings.Builder
06_structs/         → Struct definition, methods, nested structs
07_interfaces/      → Implicit interfaces, polymorphism
08_pointers/        → Pointer mechanics, value vs pointer receivers
09_goroutines/      → Goroutines, WaitGroups, concurrent fetching
10_mutexes/         → sync.Mutex, sync.RWMutex, race conditions
11_channels/        → Channel pipelines, fan-in, producers/consumers
12_generics/        → Generic functions, type constraints
13_health_checker/  → Capstone: Kubernetes-style health probe system
14_kv_store/        → Capstone: Concurrent key-value store (mini etcd)
15_worker_pool/     → Capstone: Job scheduler with worker pool
```

## Each Folder Contains

| File            | Purpose                                                        |
|-----------------|----------------------------------------------------------------|
| `exercise.go`   | Starter code with TODOs for you to fill in                    |
| `solution.go`   | Complete solution (build-tagged `ignore` — won't compile by default) |
| `README.md`     | Concept explanation, exercise description, and hints          |

## How to Work Through an Exercise

```bash
# 1. Navigate to an exercise
cd 01_variables

# 2. Read the README
cat README.md

# 3. Open exercise.go — read the comments and fill in the TODOs

# 4. Run your code
go run exercise.go

# 5. If stuck, peek at the solution (won't run by default due to build tag)
# Remove the "//go:build ignore" line to run it:
go run solution.go
```

## Prerequisites

- Go 1.21+ installed ([download](https://go.dev/dl/))
- A code editor (VS Code with the Go extension is recommended)
- Basic programming experience in any language

## Tips

- **Read the comments carefully.** Each `exercise.go` contains detailed guidance.
- **Run often.** Go's fast compiler makes the edit-run-fix cycle very tight.
- **Don't skip exercises.** Concepts build on each other — goroutines need structs,
  capstones need everything.
- **Check your work against the solution** only after giving it a real attempt.
- **Use `go vet` and `go fmt`** to keep your code idiomatic.

## Capstone Projects (13–15)

The final three exercises are mini-projects that combine multiple concepts in
contexts inspired by real Kubernetes patterns:

- **Health Checker** — Interfaces + goroutines + channels
- **KV Store** — Structs + RWMutex + interfaces (like a mini etcd)
- **Worker Pool** — Goroutines + channels + interfaces (like a job scheduler)

These are designed to be challenging. Take your time.

## Reference

Your study notes are in [`go_fundamentals.md`](go_fundamentals.md) — refer to
them whenever you need a refresher on a concept.
