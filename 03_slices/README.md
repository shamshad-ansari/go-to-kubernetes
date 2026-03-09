# 03 — Arrays and Slices

## Concept

Arrays in Go have a fixed size that is part of their type — `[3]int` and `[5]int`
are different types. In practice, you'll almost always use **slices** instead.

A slice is a dynamically-sized view into an underlying array. It has three
properties: a **pointer** to the array, a **length** (elements in use), and a
**capacity** (elements available before reallocation).

```go
s := make([]int, 3, 5)   // length=3, capacity=5
s = append(s, 4)          // length=4, capacity=5 (no reallocation)
s = append(s, 5, 6)       // length=6, capacity=10 (grew!)
```

Key things to remember:
- `append` may return a new slice if the capacity is exceeded.
- Slices are reference types — passing one to a function does NOT copy the data.
- Use `len()` for length and `cap()` for capacity.

## Exercise

Open `exercise.go`. You'll implement three functions:

1. **`FilterEven(nums []int)`** — Return a new slice containing only the even
   numbers from the input. Order must be preserved.

2. **`RemoveDuplicates(nums []int)`** — Return a new slice with duplicates
   removed, preserving the first occurrence of each value.

3. **`RotateLeft(nums []int, k int)`** — Rotate the slice left by k positions.
   For example, rotating `[1,2,3,4,5]` left by 2 gives `[3,4,5,1,2]`.

Run with: `go run exercise.go`

## Hint

- `%` (modulo) is your friend for checking even/odd and for wrapping indices.
- For `RemoveDuplicates`, a map makes an excellent "seen" tracker.
- For `RotateLeft`, think about what `k % len(nums)` does and how you can
  use slice expressions like `nums[k:]` and `nums[:k]`.
