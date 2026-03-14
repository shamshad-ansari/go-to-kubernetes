package main

import (
	"errors"
	"fmt"
	"strconv"
)

// SafeDivide divides a by b and returns the result.
// If b is zero, return 0 and an error with the message "division by zero".
//
// This is THE most common Go pattern. You'll see it hundreds of times
// in any real Go codebase.
func SafeDivide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("Divide by 0")
	}

	return a/b, nil
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
	if len(scores) == 0{
		return 0, "" , errors.New("Score Not Provided")
	}
	var sum float64
	for _ , s := range(scores) {
		if s < 0 || s > 100 {
			return 0, "", fmt.Errorf("invalid score: %.1f", s)
		}
		sum += s
	}

	avg := sum/float64(len(scores))

	var grade string
	switch {
	case avg >= 90:
		grade = "A"
	case avg >= 80:
		grade = "B"
	case avg >= 70:
		grade = "C"
	case avg >= 60:
		grade = "D"
	default:
		grade = "F"
	}

	return avg, grade, nil
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

	parsed, err := strconv.ParseFloat(input, 64)

	if err != nil {
	return 0, fmt.Errorf("failed to parse '%s': %w", input, err)
	}

	parsed = parsed * 2

	if parsed < 0 || parsed > 1000 {
		return 0, fmt.Errorf("result %.1f out of bounds [0, 1000]", parsed) 
	}

	return parsed, nil

	// What I learned
	// errors.New is for sentinel(guard erros) and fmt.Errorf is for formatting erros
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
