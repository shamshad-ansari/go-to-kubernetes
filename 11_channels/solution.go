//go:build ignore

package main

import (
	"fmt"
	"sync"
)

func Generate(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func Square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func FilterChan(in <-chan int, keep func(int) bool) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			if keep(n) {
				out <- n
			}
		}
		close(out)
	}()
	return out
}

func Merge(channels ...<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup

	for _, ch := range channels {
		wg.Add(1)
		go func(c <-chan int) {
			defer wg.Done()
			for n := range c {
				out <- n
			}
		}(ch)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func main() {
	fmt.Println("=== Pipeline: Generate → Square → Filter (even) ===")
	ch := Generate(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	squared := Square(ch)
	evens := FilterChan(squared, func(n int) bool { return n%2 == 0 })

	for val := range evens {
		fmt.Printf("  %d\n", val)
	}

	fmt.Println("\n=== Fan-In: Merge two generators ===")
	odds := Generate(1, 3, 5, 7, 9)
	tens := Generate(10, 20, 30, 40, 50)
	merged := Merge(odds, tens)

	for val := range merged {
		fmt.Printf("  %d\n", val)
	}
}
