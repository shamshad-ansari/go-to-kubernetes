//go:build ignore

package main

import "fmt"

func Filter[T any](slice []T, predicate func(T) bool) []T {
	result := []T{}
	for _, v := range slice {
		if predicate(v) {
			result = append(result, v)
		}
	}
	return result
}

func Map[T, U any](slice []T, transform func(T) U) []U {
	result := make([]U, len(slice))
	for i, v := range slice {
		result[i] = transform(v)
	}
	return result
}

func Reduce[T, U any](slice []T, initial U, fn func(U, T) U) U {
	acc := initial
	for _, v := range slice {
		acc = fn(acc, v)
	}
	return acc
}

func Contains[T comparable](slice []T, target T) bool {
	for _, v := range slice {
		if v == target {
			return true
		}
	}
	return false
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
