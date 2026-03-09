# 14 — Capstone: Concurrent Key-Value Store (Mini etcd)

## Context

In Kubernetes, **etcd** is the distributed key-value store that holds all cluster
state. Every object (pods, services, configmaps) lives in etcd. The Kubernetes
API server reads from and writes to etcd constantly, with many concurrent
clients.

This exercise builds a simplified, in-memory version of that concept: a
thread-safe key-value store that handles concurrent reads and writes. It uses
`sync.RWMutex` to allow many simultaneous readers while ensuring exclusive
write access.

## Concepts Combined

- **Structs and methods** — The store's data structure and API
- **Maps** — Underlying storage
- **Pointers and receivers** — All methods modify the store, so pointer receivers
- **RWMutex** — Concurrent-safe reads and writes
- **Interfaces** — The store satisfies a `Store` interface
- **Error handling** — Operations can fail (key not found, etc.)

## Exercise

Open `exercise.go`. You'll build a mini etcd:

1. **Define a `Store` interface** with:
   - `Get(key string) (string, error)`
   - `Set(key, value string)`
   - `Delete(key string) error`
   - `Keys() []string`
   - `Len() int`

2. **Define a `KVStore` struct** with `data map[string]string` and
   `mu sync.RWMutex`.

3. **`NewKVStore()`** — Constructor returning a `*KVStore`.

4. **Implement all Store methods** on `*KVStore`:
   - `Get` — read-lock, return error if key doesn't exist
   - `Set` — write-lock
   - `Delete` — write-lock, return error if key doesn't exist
   - `Keys` — read-lock, return sorted slice of all keys
   - `Len` — read-lock

5. **Test it** with concurrent goroutines performing mixed reads and writes.

Run with: `go run exercise.go`

## Hint

- Every method needs proper locking. Reads use `RLock()`/`RUnlock()`, writes
  use `Lock()`/`Unlock()`.
- Always `defer` the unlock immediately after locking.
- Use `sort.Strings()` to return keys in sorted order.
- For the concurrency test in main(), launch writers and readers simultaneously
  and verify the store remains consistent.
- `fmt.Errorf("key not found: %s", key)` for error messages.
