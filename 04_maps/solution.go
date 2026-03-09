//go:build ignore

package main

import (
	"fmt"
	"strings"
)

func WordFrequency(sentence string) map[string]int {
	words := strings.Fields(sentence)
	freq := make(map[string]int)
	for _, w := range words {
		freq[strings.ToLower(w)]++
	}
	return freq
}

func MergeMaps(a, b map[string]int) map[string]int {
	result := make(map[string]int)
	for k, v := range a {
		result[k] = v
	}
	for k, v := range b {
		result[k] += v
	}
	return result
}

func InvertMap(m map[string]string) map[string]string {
	result := make(map[string]string)
	for k, v := range m {
		result[v] = k
	}
	return result
}

func main() {
	fmt.Println("=== Word Frequency ===")
	freq := WordFrequency("Go is great and Go is fast and Go is fun")
	for word, count := range freq {
		fmt.Printf("  %-10s → %d\n", word, count)
	}

	fmt.Println("\n=== Merge Maps ===")
	a := map[string]int{"apples": 3, "bananas": 2}
	b := map[string]int{"bananas": 5, "oranges": 4}
	merged := MergeMaps(a, b)
	for k, v := range merged {
		fmt.Printf("  %-10s → %d\n", k, v)
	}
	fmt.Println("  Original a:", a)
	fmt.Println("  Original b:", b)

	fmt.Println("\n=== Invert Map ===")
	original := map[string]string{"name": "alice", "role": "admin", "team": "platform"}
	inverted := InvertMap(original)
	for k, v := range inverted {
		fmt.Printf("  %-10s → %s\n", k, v)
	}
}
