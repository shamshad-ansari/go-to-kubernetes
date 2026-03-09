package main

import "fmt"

// Swap takes pointers to two integers and swaps their values.
//
// After calling Swap(&a, &b), the values of a and b should be exchanged.
func Swap(a, b *int) {
	// TODO: Swap the values that a and b point to.
	// Hint: Go supports tuple assignment: *a, *b = *b, *a
}

// Counter is a simple counter that tracks an integer value.
type Counter struct {
	Value int
}

// Increment adds 1 to the counter's Value.
// Think: why must this be a pointer receiver?
func (c *Counter) Increment() {
	// TODO: Add 1 to c.Value.
}

// Reset sets Value to 0.
func (c *Counter) Reset() {
	// TODO: Set c.Value to 0.
}

// IsZero returns true if Value is 0.
// Think: why is a value receiver fine here?
func (c Counter) IsZero() bool {
	// TODO: Return true if Value == 0.
	return false // replace this
}

// ApplyToEach takes a pointer to a slice of ints and a transformation function.
// It applies the function to every element in the slice IN PLACE.
//
// Example:
//
//	nums := []int{1, 2, 3}
//	ApplyToEach(&nums, func(n int) int { return n * 2 })
//	// nums is now [2, 4, 6]
func ApplyToEach(nums *[]int, fn func(int) int) {
	// TODO: Loop through the slice (via the pointer) and apply fn to each element.
	// Remember: dereference the pointer to access the slice: (*nums)
}

func main() {
	fmt.Println("=== Swap ===")
	a, b := 10, 20
	fmt.Printf("  Before: a=%d, b=%d\n", a, b)
	Swap(&a, &b)
	fmt.Printf("  After:  a=%d, b=%d\n", a, b)

	fmt.Println("\n=== Counter ===")
	c := Counter{}
	fmt.Printf("  Initial: %d (IsZero: %t)\n", c.Value, c.IsZero())
	c.Increment()
	c.Increment()
	c.Increment()
	fmt.Printf("  After 3x Increment: %d (IsZero: %t)\n", c.Value, c.IsZero())
	c.Reset()
	fmt.Printf("  After Reset: %d (IsZero: %t)\n", c.Value, c.IsZero())

	fmt.Println("\n=== Apply To Each ===")
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println("  Before:", nums)
	ApplyToEach(&nums, func(n int) int { return n * n })
	fmt.Println("  After squaring:", nums)
	ApplyToEach(&nums, func(n int) int { return n + 10 })
	fmt.Println("  After adding 10:", nums)
}
