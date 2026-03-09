# 07 — Interfaces

## Concept

Interfaces in Go are satisfied **implicitly**. If a type has all the methods an
interface requires, it implements that interface — no `implements` keyword needed.

```go
type Shape interface {
    Area() float64
    Perimeter() float64
}
```

Any struct with `Area() float64` and `Perimeter() float64` methods automatically
satisfies `Shape`. This is incredibly powerful:

- You can define interfaces *after* the implementing types exist.
- Small interfaces (1-2 methods) are preferred — `io.Reader`, `io.Writer`,
  `error` are all single-method interfaces.
- In Kubernetes, interfaces enable swapping implementations (e.g., different
  cloud providers implementing the same storage interface).

The empty interface `interface{}` (or `any` in Go 1.18+) accepts any type,
but you lose type safety.

## Exercise

Open `exercise.go`. You'll implement a geometry toolkit:

1. **Define a `Shape` interface** with `Area() float64` and `Perimeter() float64`.

2. **Implement `Rectangle`** with `Width` and `Height` fields.

3. **Implement `Circle`** with `Radius` field. Use `math.Pi`.

4. **Implement `Triangle`** with `A`, `B`, `C` (side lengths) and `Height` fields.
   Area = 0.5 * A * Height. Perimeter = A + B + C.

5. **`LargestShape(shapes []Shape)`** — Return the shape with the largest area.

6. **`TotalArea(shapes []Shape)`** — Return the sum of all areas.

Run with: `go run exercise.go`

## Hint

- You don't need to declare that Rectangle "implements" Shape — just give it
  the right methods.
- `math.Pi` gives you pi. Import `"math"`.
- For `LargestShape`, loop through and track the one with the max area.
- Think about what happens if the slice is empty — return `nil` (interfaces can
  be nil).
