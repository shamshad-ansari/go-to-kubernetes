//go:build ignore

package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

type Triangle struct {
	SideA  float64
	SideB  float64
	SideC  float64
	Height float64
}

func (t Triangle) Area() float64 {
	return 0.5 * t.SideA * t.Height
}

func (t Triangle) Perimeter() float64 {
	return t.SideA + t.SideB + t.SideC
}

func LargestShape(shapes []Shape) Shape {
	if len(shapes) == 0 {
		return nil
	}

	largest := shapes[0]
	for _, s := range shapes[1:] {
		if s.Area() > largest.Area() {
			largest = s
		}
	}
	return largest
}

func TotalArea(shapes []Shape) float64 {
	var total float64
	for _, s := range shapes {
		total += s.Area()
	}
	return total
}

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
