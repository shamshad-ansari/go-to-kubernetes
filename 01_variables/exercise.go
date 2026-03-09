package main

import "fmt"

// ZeroValues returns a string showing the zero value of four Go types:
// int, float64, string, and bool.
//
// Expected output format: "int: 0, float64: 0, string: '', bool: false"
//
// Why this matters: Go guarantees every variable has a usable zero value.
// In Kubernetes source code, this property is relied on heavily — structs
// are often partially initialized and the zero values of unset fields
// are meaningful.
func ZeroValues() string {
	// TODO: Declare four variables (int, float64, string, bool) WITHOUT
	// assigning values to them. Use the `var` keyword.

	// TODO: Return a formatted string using fmt.Sprintf showing each
	// variable's value. Use this exact format:
	// "int: %d, float64: %.1f, string: '%s', bool: %t"

	return "" // replace this
}

// TypeConverter takes a temperature in Celsius and returns it in Fahrenheit
// as a rounded integer.
//
// Formula: F = C * 9/5 + 32
//
// Why this matters: Go does NOT do implicit type conversion. You must
// explicitly convert between numeric types. This catches subtle bugs
// that slip through in dynamically typed languages.
func TypeConverter(celsius float64) int {
	// TODO: Apply the Fahrenheit formula.
	// TODO: Convert the float64 result to int and return it.
	// Remember: int() truncates toward zero, it does not round.
	// For rounding, add 0.5 before truncating (for positive values).

	return 0 // replace this
}

// ConstantReport demonstrates Go constants.
//
// Return a string containing:
//   - The value of Pi (define as an untyped constant = 3.14159)
//   - The value of MaxRetries (define as a typed int constant = 5)
//   - The product of Pi * MaxRetries (this tests untyped constant flexibility)
//
// Format: "pi: 3.14159, maxRetries: 5, product: 15.70795"
func ConstantReport() string {
	// TODO: Declare an untyped constant Pi = 3.14159
	// TODO: Declare a typed constant MaxRetries int = 5
	// TODO: Compute the product of Pi and float64(MaxRetries)
	// TODO: Return the formatted string using fmt.Sprintf
	// Use format: "pi: %.5f, maxRetries: %d, product: %.5f"

	return "" // replace this
}

func main() {
	fmt.Println("=== Zero Values ===")
	fmt.Println(ZeroValues())

	fmt.Println("\n=== Type Converter ===")
	fmt.Printf("0°C = %d°F\n", TypeConverter(0))
	fmt.Printf("100°C = %d°F\n", TypeConverter(100))
	fmt.Printf("37°C = %d°F\n", TypeConverter(37))

	fmt.Println("\n=== Constant Report ===")
	fmt.Println(ConstantReport())
}
