//go:build ignore

package main

import "fmt"

func ZeroValues() string {
	var i int
	var f float64
	var s string
	var b bool
	return fmt.Sprintf("int: %d, float64: %.1f, string: '%s', bool: %t", i, f, s, b)
}

func TypeConverter(celsius float64) int {
	fahrenheit := celsius*9.0/5.0 + 32.0
	return int(fahrenheit + 0.5)
}

func ConstantReport() string {
	const Pi = 3.14159
	const MaxRetries int = 5
	product := Pi * float64(MaxRetries)
	return fmt.Sprintf("pi: %.5f, maxRetries: %d, product: %.5f", Pi, MaxRetries, product)
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
