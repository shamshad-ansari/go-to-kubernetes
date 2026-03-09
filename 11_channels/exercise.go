package main

import (
	"fmt"
	"sync"
)

// Generate sends all provided numbers into a channel, then closes it.
// Returns a receive-only channel.
//
// Usage: ch := Generate(1, 2, 3, 4, 5)
func Generate(nums ...int) <-chan int {
	// TODO: Create a channel.
	// TODO: Launch a goroutine that sends each number into the channel,
	//       then closes it.
	// TODO: Return the channel immediately (don't wait for the goroutine).

	return nil // replace this
}

// Square reads integers from `in`, squares each one, and sends the result
// to an output channel. Closes the output when input is exhausted.
func Square(in <-chan int) <-chan int {
	// TODO: Create an output channel.
	// TODO: Launch a goroutine that ranges over `in`, squares each value,
	//       and sends it to the output channel. Close output when done.
	// TODO: Return the output channel.

	return nil // replace this
}

// FilterChan reads from `in` and only forwards values where keep() returns true.
// Closes the output when input is exhausted.
func FilterChan(in <-chan int, keep func(int) bool) <-chan int {
	// TODO: Create an output channel.
	// TODO: Launch a goroutine that ranges over `in`. For each value,
	//       if keep(value) is true, send it to output. Close when done.
	// TODO: Return the output channel.

	return nil // replace this
}

// Merge takes multiple input channels and combines them into a single output
// channel (fan-in pattern). Values from all inputs are forwarded to the output
// concurrently. The output is closed once ALL inputs are drained.
func Merge(channels ...<-chan int) <-chan int {
	// TODO: Create an output channel.
	// TODO: Create a sync.WaitGroup.
	// TODO: For each input channel, Add(1) and launch a goroutine that:
	//       - Ranges over the input channel
	//       - Sends each value to output
	//       - Calls wg.Done() when the input channel is drained
	// TODO: Launch another goroutine that calls wg.Wait() and then
	//       closes the output channel.
	// TODO: Return the output channel.

	return nil // replace this
}

// Ensure sync is used.
var _ = sync.WaitGroup{}

func main() {
	fmt.Println("=== Pipeline: Generate → Square → Filter (even) ===")
	// Build pipeline: generate 1-10 → square each → keep only even results
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
