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
	var i int;
	var f float64;
	var s string;
	var b bool;

	return fmt.Sprintf("int: %d, float64: %.1f, string: '%s', bool: %t", i, f, s, b)
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
	f := celsius * 9.0/5.0 + 32.0
	return int(f+0.5)

	// What I learned
	// No automatic type conversion, so need to explicity convert before you can 
	// do any airthimetic operations
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
	const Pi = 3.14159
	const MaxRetries int = 5
	product := Pi * float64(MaxRetries)

	return fmt.Sprintf("pi: %.5f, maxRetries: %d, product: %.5f ", Pi, MaxRetries, product)
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
