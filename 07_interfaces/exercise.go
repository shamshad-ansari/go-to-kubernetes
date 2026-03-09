package main

import (
	"fmt"
	"math"
)

// Shape is an interface that any geometry must implement.
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Rectangle has a Width and Height.
type Rectangle struct {
	Width  float64
	Height float64
}

// Area returns the area of the rectangle: Width * Height.
func (r Rectangle) Area() float64 {
	// TODO: Return Width * Height.
	return 0 // replace this
}

// Perimeter returns the perimeter: 2 * (Width + Height).
func (r Rectangle) Perimeter() float64 {
	// TODO: Return 2 * (Width + Height).
	return 0 // replace this
}

// Circle has a Radius.
type Circle struct {
	Radius float64
}

// Area returns the area: Pi * Radius^2.
func (c Circle) Area() float64 {
	// TODO: Return math.Pi * Radius * Radius.
	return 0 // replace this
}

// Perimeter returns the circumference: 2 * Pi * Radius.
func (c Circle) Perimeter() float64 {
	// TODO: Return 2 * math.Pi * Radius.
	return 0 // replace this
}

// Triangle has three sides and a height (relative to SideA).
type Triangle struct {
	SideA  float64
	SideB  float64
	SideC  float64
	Height float64
}

// Area returns the area: 0.5 * SideA * Height.
func (t Triangle) Area() float64 {
	// TODO: Return 0.5 * SideA * Height.
	return 0 // replace this
}

// Perimeter returns SideA + SideB + SideC.
func (t Triangle) Perimeter() float64 {
	// TODO: Return the sum of all three sides.
	return 0 // replace this
}

// LargestShape returns the shape with the largest area from the slice.
// If the slice is empty, return nil.
func LargestShape(shapes []Shape) Shape {
	// TODO: Handle empty slice.
	// TODO: Loop through shapes, tracking the one with maximum area.
	// TODO: Return the largest.

	return nil // replace this
}

// TotalArea returns the sum of areas of all shapes.
func TotalArea(shapes []Shape) float64 {
	// TODO: Sum up Area() for every shape.

	return 0 // replace this
}

// Ensure math is used (remove this line once you use math.Pi in your methods).
var _ = math.Pi

func main() {
	shapes := []Shape{
		Rectangle{Width: 10, Height: 5},
		Circle{Radius: 7},
		Triangle{SideA: 6, SideB: 8, SideC: 10, Height: 4.8},
	}

	fmt.Println("=== Shape Details ===")
	for _, s := range shapes {
		fmt.Printf("  %-12T → Area: %8.2f, Perimeter: %8.2f\n", s, s.Area(), s.Perimeter())
	}

	fmt.Println("\n=== Analysis ===")
	largest := LargestShape(shapes)
	if largest != nil {
		fmt.Printf("  Largest shape: %T with area %.2f\n", largest, largest.Area())
	}
	fmt.Printf("  Total area:    %.2f\n", TotalArea(shapes))
}
