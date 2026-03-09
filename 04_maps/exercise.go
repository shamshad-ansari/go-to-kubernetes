package main

import (
	"fmt"
)

// WordFrequency takes a sentence and returns a map of each word to its count.
// Words should be compared case-insensitively (convert to lowercase).
//
// Example: WordFrequency("Go is great and Go is fast") →
//
//	map[string]int{"go": 2, "is": 2, "great": 1, "and": 1, "fast": 1}
func WordFrequency(sentence string) map[string]int {
	// TODO: Split the sentence into words.
	//       Hint: strings.Fields() splits on whitespace.
	// TODO: Create an empty map[string]int.
	// TODO: Loop through words. Lowercase each word with strings.ToLower().
	// TODO: Increment the count for that word in the map.
	// TODO: Return the map.

	return nil // replace this
}

// MergeMaps merges two maps into a new one. If both maps contain the same key,
// their values are summed. The original maps must not be modified.
//
// Example: MergeMaps({"a": 1, "b": 2}, {"b": 3, "c": 4}) →
//
//	{"a": 1, "b": 5, "c": 4}
func MergeMaps(a, b map[string]int) map[string]int {
	// TODO: Create a new result map.
	// TODO: Copy all entries from a into result.
	// TODO: Loop through b. For each key, add its value to what's in result.
	//       (If the key doesn't exist in result, the zero value 0 is used.)
	// TODO: Return result.

	return nil // replace this
}

// InvertMap swaps keys and values in the given map.
// If multiple keys map to the same value, any one of them can be the key
// in the result (map iteration is random, so this is non-deterministic).
//
// Example: InvertMap({"a": "x", "b": "y"}) → {"x": "a", "y": "b"}
func InvertMap(m map[string]string) map[string]string {
	// TODO: Create a new result map.
	// TODO: Loop through m. For each key-value pair, add the inverted entry
	//       (value becomes key, key becomes value).
	// TODO: Return result.

	return nil // replace this
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
