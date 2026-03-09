# 08 — Pointers

## Concept

A pointer holds the memory address of a variable. Go uses pointers to let
functions modify data without copying it.

```go
x := 42
p := &x     // p points to x
*p = 100    // x is now 100
```

| Syntax | Meaning |
|--------|---------|
| `*T`   | Type: "pointer to T" |
| `&x`   | "address of x" |
| `*p`   | "value at the address p holds" (dereference) |

**Value receiver** `func (v Vehicle) Method()` — gets a copy, can't modify the
original.

**Pointer receiver** `func (v *Vehicle) Method()` — gets the address, can modify
the original.

Rule of thumb in Go:
- Use pointer receivers when the method needs to modify the receiver.
- Use pointer receivers for large structs to avoid copying.
- Be consistent — if one method uses a pointer receiver, all should.

Slices, maps, and channels are already reference types — you don't need pointers
for them.

## Exercise

Open `exercise.go`. You'll work with pointers and receivers:

1. **`Swap(a, b *int)`** — Swap the values of two integers using pointers.

2. **`Counter` struct** with a `Value int` field:
   - `Increment()` — add 1 to Value (needs pointer receiver to modify)
   - `Reset()` — set Value to 0 (needs pointer receiver)
   - `IsZero()` — return true if Value is 0 (value receiver is fine)

3. **`ApplyToEach(nums *[]int, fn func(int) int)`** — Apply a transformation
   function to every element of a slice in-place via a pointer.

Run with: `go run exercise.go`

## Hint

- `Swap`: use a temporary variable, or Go's tuple assignment `*a, *b = *b, *a`.
- For `Counter` methods, if you use `func (c Counter) Increment()`, the change
  won't persist — the receiver is a copy. Use `func (c *Counter) Increment()`.
- `ApplyToEach` receives `*[]int`. Dereference with `(*nums)[i]` to access
  elements.
