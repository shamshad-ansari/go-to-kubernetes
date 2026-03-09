//go:build ignore

package main

import (
	"fmt"
	"sort"
	"sync"
)

type Scoreboard struct {
	scores map[string]int
	mu     sync.RWMutex
}

func NewScoreboard() *Scoreboard {
	return &Scoreboard{
		scores: make(map[string]int),
	}
}

func (sb *Scoreboard) RecordScore(name string, points int) {
	sb.mu.Lock()
	defer sb.mu.Unlock()
	sb.scores[name] += points
}

func (sb *Scoreboard) GetScore(name string) (int, bool) {
	sb.mu.RLock()
	defer sb.mu.RUnlock()
	score, ok := sb.scores[name]
	return score, ok
}

func (sb *Scoreboard) TopScorers(n int) []string {
	sb.mu.RLock()
	defer sb.mu.RUnlock()

	names := make([]string, 0, len(sb.scores))
	for name := range sb.scores {
		names = append(names, name)
	}

	sort.Slice(names, func(i, j int) bool {
		return sb.scores[names[i]] > sb.scores[names[j]]
	})

	if n > len(names) {
		n = len(names)
	}
	return names[:n]
}

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
