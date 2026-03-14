package main

import "fmt"

// FilterEven returns a new slice containing only the even numbers from nums.
// The order of elements must be preserved.
//
// Example: FilterEven([]int{1, 2, 3, 4, 5, 6}) → [2, 4, 6]
func FilterEven(nums []int) []int {
	result := []int{}
	for _, item := range(nums){
		if (item % 2 == 0) {
			result = append(result, item)
		}
	}
	return result
}

// RemoveDuplicates returns a new slice with duplicate values removed.
// The first occurrence of each value is kept, and order is preserved.
//
// Example: RemoveDuplicates([]int{3, 1, 4, 1, 5, 9, 2, 6, 5}) → [3, 1, 4, 5, 9, 2, 6]
func RemoveDuplicates(nums []int) []int {
	seen := make(map[int]bool)
	result := []int{}
	for _, item := range(nums){
		if !seen[item]{
			seen[item] = true
			result = append(result, item)
		}
		
	}
	return result
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
	n := len(nums)
	if n <= 1 {
		return nums
	}

	k %= n
	if k < 0 {
		k += n
	}

	result := make([]int, n)
	copy(result, nums[k:])
	copy(result[n-k:], nums[:k])

	return result

	// What I learned
	// copy here writes into a slice of result (which is itself a slice). 
	// Slices share the same underlying array, so updating any slice of result 
	// updates the underlying array, and the changes are reflected in result.
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
