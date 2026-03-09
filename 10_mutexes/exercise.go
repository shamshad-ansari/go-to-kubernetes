package main

import (
	"fmt"
	"sort"
	"sync"
)

// Scoreboard tracks player scores with thread-safe access.
type Scoreboard struct {
	scores map[string]int
	mu     sync.RWMutex
}

// NewScoreboard creates and returns a pointer to a new Scoreboard
// with an initialized map.
func NewScoreboard() *Scoreboard {
	// TODO: Return a &Scoreboard with an initialized map.

	return nil // replace this
}

// RecordScore adds `points` to the player's score.
// If the player doesn't exist yet, their score starts from 0 (map zero value).
// This method MUST be safe for concurrent use (write lock).
func (sb *Scoreboard) RecordScore(name string, points int) {
	// TODO: Lock the mutex (write lock — not RLock).
	// TODO: Defer Unlock.
	// TODO: Add points to sb.scores[name].
}

// GetScore returns the score for a player and whether they exist.
// This method MUST be safe for concurrent use (read lock).
func (sb *Scoreboard) GetScore(name string) (int, bool) {
	// TODO: RLock the mutex.
	// TODO: Defer RUnlock.
	// TODO: Look up the name in the map using the two-value form.
	// TODO: Return score and existence boolean.

	return 0, false // replace this
}

// TopScorers returns the names of the top N players, sorted by score
// descending. If there are fewer than N players, return all of them.
// This method MUST be safe for concurrent use (read lock).
func (sb *Scoreboard) TopScorers(n int) []string {
	// TODO: RLock the mutex.
	// TODO: Defer RUnlock.
	// TODO: Collect all player names into a slice.
	// TODO: Sort the slice by score (descending) using sort.Slice.
	// TODO: Return the top N (or all if fewer than N exist).

	return nil // replace this
}

// Ensure sort is used.
var _ = sort.Slice

func main() {
	sb := NewScoreboard()

	players := []string{"alice", "bob", "carol", "dave", "eve"}

	fmt.Println("=== Concurrent Score Recording ===")
	var wg sync.WaitGroup

	for _, player := range players {
		for i := 0; i < 100; i++ {
			wg.Add(1)
			go func(name string, points int) {
				defer wg.Done()
				sb.RecordScore(name, points)
			}(player, i%10+1)
		}
	}

	wg.Wait()

	fmt.Println("\n=== Final Scores ===")
	for _, player := range players {
		score, _ := sb.GetScore(player)
		fmt.Printf("  %-8s → %d points\n", player, score)
	}

	fmt.Println("\n=== Top 3 Scorers ===")
	top := sb.TopScorers(3)
	for i, name := range top {
		score, _ := sb.GetScore(name)
		fmt.Printf("  #%d: %-8s → %d points\n", i+1, name, score)
	}

	_, exists := sb.GetScore("nobody")
	fmt.Printf("\n  'nobody' exists: %t\n", exists)
}
