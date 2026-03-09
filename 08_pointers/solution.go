//go:build ignore

package main

import "fmt"

func Swap(a, b *int) {
	*a, *b = *b, *a
}

type Counter struct {
	Value int
}

func (c *Counter) Increment() {
	c.Value++
}

func (c *Counter) Reset() {
	c.Value = 0
}

func (c Counter) IsZero() bool {
	return c.Value == 0
}

func ApplyToEach(nums *[]int, fn func(int) int) {
	for i, v := range *nums {
		(*nums)[i] = fn(v)
	}
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
