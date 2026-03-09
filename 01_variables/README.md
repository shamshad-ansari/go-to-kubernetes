# 01 — Variables and Types

## Concept

Go is statically and strongly typed. Every variable has a type known at compile
time, and implicit type mixing is not allowed.

There are three ways to declare variables:

```go
var x int = 10       // explicit type
var y = 10           // type inferred
z := 10              // short declaration (most common inside functions)
```

Every type has a **zero value** — the default when no value is assigned:
- `int` → `0`, `float64` → `0.0`, `string` → `""`, `bool` → `false`

Constants are declared with `const` and must have a value at declaration time.

## Exercise

Open `exercise.go`. You'll implement three functions:

1. **`ZeroValues()`** — Return the zero values of four different types as a
   formatted string. This tests your knowledge of Go's zero value guarantee.

2. **`TypeConverter(celsius float64)`** — Convert a Celsius temperature to
   Fahrenheit and return it as a rounded integer. This tests explicit type
   conversion (Go won't do it implicitly).

3. **`ConstantReport()`** — Work with constants and demonstrate the difference
   between typed and untyped constants.

Run with: `go run exercise.go`

## Hint

- Use `fmt.Sprintf` to build formatted strings.
- To convert `float64` to `int`, use `int(value)` — Go truncates toward zero.
- Untyped constants in Go are flexible — they adapt to the context where they're
  used. Typed constants behave like variables of that type.
