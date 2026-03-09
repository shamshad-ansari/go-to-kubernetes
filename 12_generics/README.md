# 12 — Generics

## Concept

Generics (introduced in Go 1.18) let you write functions and types that work
with multiple types while keeping full type safety.

```go
func Contains[T comparable](slice []T, target T) bool {
    for _, v := range slice {
        if v == target {
            return true
        }
    }
    return false
}
```

Type constraints define what types are allowed:
- `any` — any type at all (alias for `interface{}`)
- `comparable` — types that support `==` and `!=`
- Custom constraints using interfaces with type lists

```go
type Number interface {
    int | int32 | int64 | float32 | float64
}
```

Before generics, Go developers had to write separate functions for each type or
use `interface{}` and lose type safety. Generics solve this elegantly.

In Kubernetes, generics are increasingly used in newer utilities and libraries
(e.g., `k8s.io/utils`).

## Exercise

Open `exercise.go`. You'll build a generic utility toolkit:

1. **`Filter[T any](slice []T, predicate func(T) bool) []T`** — Return a new
   slice with only elements where predicate returns true.

2. **`Map[T, U any](slice []T, transform func(T) U) []U`** — Transform each
   element from type T to type U.

3. **`Reduce[T, U any](slice []T, initial U, fn func(U, T) U) U`** — Fold a
   slice into a single value using an accumulator.

4. **`Contains[T comparable](slice []T, target T) bool`** — Check if a value
   exists in a slice.

5. **Use all four** in main() to demonstrate they work with different types
   (ints, strings, structs).

Run with: `go run exercise.go`

## Hint

- The type parameter goes in square brackets before the regular parameters:
  `func Name[T constraint](param T)`.
- `any` allows any type but you can't do much with it (no operators).
- `comparable` lets you use `==` and `!=`.
- `Filter` and `Map` are similar in structure — loop, apply function, collect.
- `Reduce` carries an accumulator through the loop.
