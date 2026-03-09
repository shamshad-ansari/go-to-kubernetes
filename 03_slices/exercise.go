package main

import "fmt"

// FilterEven returns a new slice containing only the even numbers from nums.
// The order of elements must be preserved.
//
// Example: FilterEven([]int{1, 2, 3, 4, 5, 6}) → [2, 4, 6]
func FilterEven(nums []int) []int {
	// TODO: Create an empty result slice.
	// TODO: Loop through nums. If a number is even (num%2 == 0), append it.
	// TODO: Return the result.

	return nil // replace this
}

// RemoveDuplicates returns a new slice with duplicate values removed.
// The first occurrence of each value is kept, and order is preserved.
//
// Example: RemoveDuplicates([]int{3, 1, 4, 1, 5, 9, 2, 6, 5}) → [3, 1, 4, 5, 9, 2, 6]
func RemoveDuplicates(nums []int) []int {
	// TODO: Create a map to track which numbers you've already seen.
	// TODO: Create an empty result slice.
	// TODO: Loop through nums. If the number hasn't been seen, add it to
	//       result and mark it as seen.
	// TODO: Return the result.

	return nil // replace this
}

// RotateLeft rotates a slice to the left by k positions.
// Elements that fall off the left end wrap around to the right.
//
// Example: RotateLeft([]int{1, 2, 3, 4, 5}, 2) → [3, 4, 5, 1, 2]
//
// Handle edge cases:
//   - If the slice is empty or has one element, return it as-is.
//   - If k is larger than the slice length, it should wrap (use modulo).
//   - If k is negative, treat it as a right rotation.
func RotateLeft(nums []int, k int) []int {
	// TODO: Handle edge cases (empty slice, single element).
	// TODO: Normalize k using modulo so it's within [0, len(nums)).
	//       Hint: k = k % len(nums); handle negative k.
	// TODO: Build the rotated slice using slice expressions.
	//       The result is nums[k:] followed by nums[:k].

	return nil // replace this
}

func main() {
	fmt.Println("=== Filter Even ===")
	fmt.Println(FilterEven([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}))
	fmt.Println(FilterEven([]int{1, 3, 5, 7}))
	fmt.Println(FilterEven([]int{}))

	fmt.Println("\n=== Remove Duplicates ===")
	fmt.Println(RemoveDuplicates([]int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3}))
	fmt.Println(RemoveDuplicates([]int{1, 1, 1, 1}))
	fmt.Println(RemoveDuplicates([]int{5}))

	fmt.Println("\n=== Rotate Left ===")
	fmt.Println(RotateLeft([]int{1, 2, 3, 4, 5}, 2))
	fmt.Println(RotateLeft([]int{1, 2, 3, 4, 5}, 7))
	fmt.Println(RotateLeft([]int{1, 2, 3}, -1))
	fmt.Println(RotateLeft([]int{}, 3))
}
