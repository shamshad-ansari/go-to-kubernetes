//go:build ignore

package main

import (
	"errors"
	"fmt"
	"strconv"
)

func SafeDivide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

func GradeCalculator(scores []float64) (float64, string, error) {
	if len(scores) == 0 {
		return 0, "", errors.New("no scores provided")
	}

	var sum float64
	for _, s := range scores {
		if s < 0 || s > 100 {
			return 0, "", fmt.Errorf("invalid score: %.1f", s)
		}
		sum += s
	}

	avg := sum / float64(len(scores))

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

func ChainedOperation(input string) (float64, error) {
	parsed, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return 0, fmt.Errorf("failed to parse '%s': %w", input, err)
	}

	doubled := parsed * 2

	if doubled < 0 || doubled > 1000 {
		return 0, fmt.Errorf("result %.1f out of bounds [0, 1000]", doubled)
	}

	return doubled, nil
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
