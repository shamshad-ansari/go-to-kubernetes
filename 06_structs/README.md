# 06 — Structs and Methods

## Concept

Structs are Go's way of defining custom data types by grouping fields together.
Unlike classes in OOP languages, Go structs have **no inheritance**. Instead, you
compose behavior by embedding structs and attaching methods.

```go
type Server struct {
    Name   string
    Port   int
    Region string
}

func (s Server) Address() string {
    return fmt.Sprintf("%s:%d", s.Name, s.Port)
}
```

The `(s Server)` before the function name is called a **receiver**. It attaches
the method to the `Server` type. This is how Go achieves method dispatch without
classes.

Structs can be nested — a struct field can itself be a struct. This is used
constantly in Kubernetes API objects (a Pod contains a Spec, which contains
Containers, each of which contains Resources...).

## Exercise

Open `exercise.go`. You'll build a small fleet management system:

1. **Define an `Engine` struct** with `HorsePower int` and `FuelCapacity float64`.

2. **Define a `Vehicle` struct** with `Make string`, `Model string`,
   `Year int`, `Mileage float64`, `MPG float64`, and a nested `Engine`.

3. **`FuelRange()` method** — Returns how many miles the vehicle can travel on
   a full tank (`Engine.FuelCapacity * MPG`).

4. **`NeedsService()` method** — Returns true if mileage exceeds 100,000.

5. **`FleetSummary(vehicles []Vehicle)` function** — Returns a formatted string
   summarizing the fleet: total vehicles, average fuel range, and count of
   vehicles needing service.

Run with: `go run exercise.go`

## Hint

- Methods are defined outside the struct, with a receiver: `func (v Vehicle) MethodName()`.
- Access nested struct fields with dot notation: `v.Engine.FuelCapacity`.
- `fmt.Sprintf` is your friend for building summary strings.
