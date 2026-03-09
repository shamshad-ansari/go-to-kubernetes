//go:build ignore

package main

import "fmt"

func FilterEven(nums []int) []int {
	result := []int{}
	for _, n := range nums {
		if n%2 == 0 {
			result = append(result, n)
		}
	}
	return result
}

func RemoveDuplicates(nums []int) []int {
	seen := make(map[int]bool)
	result := []int{}
	for _, n := range nums {
		if !seen[n] {
			seen[n] = true
			result = append(result, n)
		}
	}
	return result
}

func RotateLeft(nums []int, k int) []int {
	if len(nums) <= 1 {
		return nums
	}

	n := len(nums)
	k = k % n
	if k < 0 {
		k += n
	}

	result := make([]int, n)
	copy(result, nums[k:])
	copy(result[n-k:], nums[:k])
	return result
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
