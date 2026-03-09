package main

import (
	"fmt"
	"sort"
	"sync"
)

// Store defines the interface for a key-value store.
type Store interface {
	Get(key string) (string, error)
	Set(key, value string)
	Delete(key string) error
	Keys() []string
	Len() int
}

// KVStore is a concurrent-safe in-memory key-value store.
type KVStore struct {
	data map[string]string
	mu   sync.RWMutex
}

// NewKVStore creates and returns a new KVStore with an initialized map.
func NewKVStore() *KVStore {
	// TODO: Return a new KVStore with initialized data map.

	return nil // replace this
}

// Get retrieves the value for a key. Returns an error if the key doesn't exist.
// Must use a read lock.
func (kv *KVStore) Get(key string) (string, error) {
	// TODO: RLock/RUnlock.
	// TODO: Look up key with the two-value form.
	// TODO: If not found, return error: "key not found: <key>"
	// TODO: Return the value and nil.

	return "", nil // replace this
}

// Set stores a key-value pair. If the key already exists, it's overwritten.
// Must use a write lock.
func (kv *KVStore) Set(key, value string) {
	// TODO: Lock/Unlock.
	// TODO: Set the key in the map.
}

// Delete removes a key from the store. Returns an error if the key doesn't exist.
// Must use a write lock.
func (kv *KVStore) Delete(key string) error {
	// TODO: Lock/Unlock.
	// TODO: Check if key exists. If not, return error: "key not found: <key>"
	// TODO: Delete the key.

	return nil // replace this
}

// Keys returns a sorted slice of all keys in the store.
// Must use a read lock.
func (kv *KVStore) Keys() []string {
	// TODO: RLock/RUnlock.
	// TODO: Collect all keys from the map.
	// TODO: Sort them with sort.Strings().
	// TODO: Return the sorted slice.

	return nil // replace this
}

// Len returns the number of entries in the store.
// Must use a read lock.
func (kv *KVStore) Len() int {
	// TODO: RLock/RUnlock.
	// TODO: Return len(kv.data).

	return 0 // replace this
}

// Ensure imports are used.
var _ = sort.Strings

func main() {
	store := NewKVStore()

	fmt.Println("=== Basic Operations ===")
	store.Set("/registry/pods/nginx", `{"name": "nginx", "status": "running"}`)
	store.Set("/registry/pods/redis", `{"name": "redis", "status": "pending"}`)
	store.Set("/registry/services/web", `{"name": "web", "port": 80}`)

	val, err := store.Get("/registry/pods/nginx")
	if err != nil {
		fmt.Printf("  Error: %v\n", err)
	} else {
		fmt.Printf("  GET /registry/pods/nginx → %s\n", val)
	}

	_, err = store.Get("/registry/pods/unknown")
	if err != nil {
		fmt.Printf("  GET /registry/pods/unknown → Error: %v\n", err)
	}

	fmt.Printf("  Keys: %v\n", store.Keys())
	fmt.Printf("  Len: %d\n", store.Len())

	err = store.Delete("/registry/pods/redis")
	fmt.Printf("  DELETE /registry/pods/redis → err: %v\n", err)
	fmt.Printf("  Keys after delete: %v\n", store.Keys())

	err = store.Delete("/registry/pods/redis")
	if err != nil {
		fmt.Printf("  DELETE again → Error: %v\n", err)
	}

	fmt.Println("\n=== Concurrent Access Test ===")
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			key := fmt.Sprintf("/registry/pods/pod-%d", id)
			value := fmt.Sprintf(`{"id": %d}`, id)
			store.Set(key, value)
		}(i)
	}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			key := fmt.Sprintf("/registry/pods/pod-%d", id)
			store.Get(key)
		}(i)
	}

	wg.Wait()

	fmt.Printf("  Store size after concurrent writes: %d\n", store.Len())
	fmt.Printf("  First 5 keys: %v\n", store.Keys()[:5])

	// Verify the store interface is satisfied.
	var _ Store = store
	fmt.Println("\n  KVStore satisfies the Store interface.")
}
