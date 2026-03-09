# 02 — Functions and Error Handling

## Concept

Go functions can return multiple values. The most important pattern in all of Go
is returning an `error` as the last return value:

```go
result, err := doSomething()
if err != nil {
    // handle error
}
```

Errors in Go are **values**, not exceptions. You check them explicitly after
every call. This is a deliberate design choice — it makes error handling visible
and forces you to think about failure cases.

The `errors.New()` function creates a simple error. For formatted error messages,
use `fmt.Errorf()`.

## Exercise

Open `exercise.go`. You'll implement three functions:

1. **`SafeDivide(a, b float64)`** — Divide two numbers, returning an error if
   the denominator is zero. This is the classic Go error pattern.

2. **`GradeCalculator(scores []float64)`** — Calculate the average of a slice
   of scores and return both the average and a letter grade. Handle edge cases:
   empty slice, scores out of range.

3. **`ChainedOperation(input string)`** — Parse a string to a number, double it,
   then check if it's within bounds. Each step can fail. This tests error
   propagation — how errors flow through a call chain.

Run with: `go run exercise.go`

## Hint

- `errors.New("message")` creates a basic error value.
- `fmt.Errorf("format %v", val)` creates a formatted error.
- `strconv.ParseFloat(s, 64)` parses a string to float64 and returns an error
  if the string isn't a valid number.
- Always check errors immediately — don't defer error checking.
