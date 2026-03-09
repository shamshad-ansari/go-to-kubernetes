//go:build ignore

package main

import (
	"fmt"
	"sort"
	"sync"
)

type Store interface {
	Get(key string) (string, error)
	Set(key, value string)
	Delete(key string) error
	Keys() []string
	Len() int
}

type KVStore struct {
	data map[string]string
	mu   sync.RWMutex
}

func NewKVStore() *KVStore {
	return &KVStore{
		data: make(map[string]string),
	}
}

func (kv *KVStore) Get(key string) (string, error) {
	kv.mu.RLock()
	defer kv.mu.RUnlock()

	val, ok := kv.data[key]
	if !ok {
		return "", fmt.Errorf("key not found: %s", key)
	}
	return val, nil
}

func (kv *KVStore) Set(key, value string) {
	kv.mu.Lock()
	defer kv.mu.Unlock()
	kv.data[key] = value
}

func (kv *KVStore) Delete(key string) error {
	kv.mu.Lock()
	defer kv.mu.Unlock()

	_, ok := kv.data[key]
	if !ok {
		return fmt.Errorf("key not found: %s", key)
	}
	delete(kv.data, key)
	return nil
}

func (kv *KVStore) Keys() []string {
	kv.mu.RLock()
	defer kv.mu.RUnlock()

	keys := make([]string, 0, len(kv.data))
	for k := range kv.data {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

func (kv *KVStore) Len() int {
	kv.mu.RLock()
	defer kv.mu.RUnlock()
	return len(kv.data)
}

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

	var _ Store = store
	fmt.Println("\n  KVStore satisfies the Store interface.")
}
