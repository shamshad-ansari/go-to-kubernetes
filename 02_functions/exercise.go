package main

import (
	"fmt"
)

// SafeDivide divides a by b and returns the result.
// If b is zero, return 0 and an error with the message "division by zero".
//
// This is THE most common Go pattern. You'll see it hundreds of times
// in any real Go codebase.
func SafeDivide(a, b float64) (float64, error) {
	// TODO: Check if b is zero. If so, return 0 and an error.
	// TODO: Otherwise, return a/b and nil.

	return 0, nil // replace this
}

// GradeCalculator takes a slice of scores (0-100) and returns:
//   - The average score (float64)
//   - A letter grade (string): A (>=90), B (>=80), C (>=70), D (>=60), F (<60)
//   - An error if the slice is empty or any score is outside 0-100
//
// Error messages:
//   - Empty slice: "no scores provided"
//   - Invalid score: "invalid score: X.X" (where X.X is the bad score)
func GradeCalculator(scores []float64) (float64, string, error) {
	// TODO: Check for empty slice, return error.
	// TODO: Validate all scores are in [0, 100], return error on first invalid.
	// TODO: Calculate the average.
	// TODO: Determine the letter grade based on the average.
	// TODO: Return average, grade, nil.

	return 0, "", nil // replace this
}

// ChainedOperation takes a numeric string, parses it to a float, doubles it,
// and checks that the result is between 0 and 1000 (inclusive).
//
// Returns:
//   - The doubled value
//   - An error if parsing fails or the result is out of bounds
//
// Error messages:
//   - Parse failure: "failed to parse 'X': <original error>"
//   - Out of bounds: "result Y out of bounds [0, 1000]"
//
// Hint: use strconv.ParseFloat and fmt.Errorf to wrap errors.
func ChainedOperation(input string) (float64, error) {
	// TODO: Parse the input string to float64 using strconv.ParseFloat.
	// TODO: If parsing fails, return a descriptive error wrapping the original.
	// TODO: Double the parsed value.
	// TODO: Check bounds [0, 1000]. Return an error if out of range.
	// TODO: Return the doubled value and nil.

	return 0, nil // replace this
}

func main() {
	fmt.Println("=== Safe Divide ===")
	result, err := SafeDivide(10, 3)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("10 / 3 = %.2f\n", result)
	}

	result, err = SafeDivide(10, 0)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("10 / 0 = %.2f\n", result)
	}

	fmt.Println("\n=== Grade Calculator ===")
	avg, grade, err := GradeCalculator([]float64{92, 85, 78, 95, 88})
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Average: %.1f, Grade: %s\n", avg, grade)
	}

	_, _, err = GradeCalculator([]float64{})
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	fmt.Println("\n=== Chained Operation ===")
	val, err := ChainedOperation("42.5")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Result: %.1f\n", val)
	}

	val, err = ChainedOperation("not_a_number")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	val, err = ChainedOperation("999")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Result: %.1f\n", val)
	}
}
