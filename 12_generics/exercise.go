package main

import "fmt"

// Filter returns a new slice containing only the elements from `slice`
// where `predicate` returns true.
//
// Example:
//
//	Filter([]int{1,2,3,4,5}, func(n int) bool { return n > 3 }) → [4, 5]
//	Filter([]string{"go","java","rust"}, func(s string) bool { return len(s) <= 3 }) → ["go"]
func Filter[T any](slice []T, predicate func(T) bool) []T {
	// TODO: Create an empty result slice.
	// TODO: Range over the input slice.
	// TODO: If predicate(element) is true, append to result.
	// TODO: Return result.

	return nil // replace this
}

// Map transforms each element of `slice` from type T to type U using
// the `transform` function.
//
// Example:
//
//	Map([]int{1,2,3}, func(n int) string { return fmt.Sprintf("#%d", n) }) → ["#1", "#2", "#3"]
func Map[T, U any](slice []T, transform func(T) U) []U {
	// TODO: Create a result slice with the same length as input.
	// TODO: Range over the input, apply transform, store in result.
	// TODO: Return result.

	return nil // replace this
}

// Reduce folds a slice into a single value. Starting from `initial`, it
// applies `fn(accumulator, element)` for each element, carrying the result
// forward.
//
// Example:
//
//	Reduce([]int{1,2,3,4}, 0, func(acc, n int) int { return acc + n }) → 10
//	Reduce([]string{"a","b","c"}, "", func(acc, s string) string { return acc + s }) → "abc"
func Reduce[T, U any](slice []T, initial U, fn func(U, T) U) U {
	// TODO: Start with `initial` as the accumulator.
	// TODO: For each element, set accumulator = fn(accumulator, element).
	// TODO: Return the final accumulator.

	var zero U
	return zero // replace this
}

// Contains checks if `target` exists in `slice`.
// The type constraint `comparable` is needed because we use == for comparison.
//
// Example:
//
//	Contains([]int{1,2,3}, 2)         → true
//	Contains([]string{"a","b"}, "c")  → false
func Contains[T comparable](slice []T, target T) bool {
	// TODO: Range over the slice.
	// TODO: If any element == target, return true.
	// TODO: Return false if not found.

	return false // replace this
}

func main() {
	fmt.Println("=== Filter ===")
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	evens := Filter(nums, func(n int) bool { return n%2 == 0 })
	fmt.Println("  Even numbers:", evens)

	words := []string{"go", "kubernetes", "docker", "api", "helm"}
	short := Filter(words, func(s string) bool { return len(s) <= 4 })
	fmt.Println("  Short words:", short)

	fmt.Println("\n=== Map ===")
	doubled := Map(nums, func(n int) int { return n * 2 })
	fmt.Println("  Doubled:", doubled)

	lengths := Map(words, func(s string) int { return len(s) })
	fmt.Println("  Word lengths:", lengths)

	labels := Map(nums[:5], func(n int) string {
		return fmt.Sprintf("item-%d", n)
	})
	fmt.Println("  Labels:", labels)

	fmt.Println("\n=== Reduce ===")
	sum := Reduce(nums, 0, func(acc, n int) int { return acc + n })
	fmt.Println("  Sum of 1-10:", sum)

	joined := Reduce(words, "", func(acc, s string) string {
		if acc == "" {
			return s
		}
		return acc + ", " + s
	})
	fmt.Println("  Joined words:", joined)

	fmt.Println("\n=== Contains ===")
	fmt.Println("  nums contains 5:", Contains(nums, 5))
	fmt.Println("  nums contains 99:", Contains(nums, 99))
	fmt.Println("  words contains 'helm':", Contains(words, "helm"))
	fmt.Println("  words contains 'python':", Contains(words, "python"))
}
