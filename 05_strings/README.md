# 05 — Strings and Runes

## Concept

Go strings are immutable sequences of **bytes**, not characters. They are encoded
in UTF-8, which means multi-byte characters like `é`, `日`, or emojis occupy
more than one byte.

This has real consequences:
- `len("résumé")` returns **8** (bytes), not 6 (characters).
- `s[1]` gives you a **byte**, not the character `é`.
- To work with characters correctly, convert to `[]rune` or use `range`.

A `rune` is an alias for `int32` and represents a single Unicode code point.

```go
s := "résumé"
runes := []rune(s)
fmt.Println(len(runes))        // 6 (characters)
fmt.Println(string(runes[1]))  // é
```

For building strings efficiently, use `strings.Builder` instead of `+`
concatenation (which creates a new string each time).

## Exercise

Open `exercise.go`. You'll implement three functions:

1. **`ReverseString(s string)`** — Reverse a string correctly, even when it
   contains multi-byte Unicode characters like `"résumé"` or `"Hello, 世界"`.

2. **`IsPalindrome(s string)`** — Check if a string reads the same forwards and
   backwards, ignoring case. Must handle Unicode correctly.

3. **`CensorWord(text, word string)`** — Replace every occurrence of `word` in
   `text` with asterisks of the same rune-length. Case-sensitive matching.

Run with: `go run exercise.go`

## Hint

- Convert to `[]rune` before reversing — this handles multi-byte chars correctly.
- `strings.ToLower()` for case-insensitive comparison.
- `strings.ReplaceAll()` can handle `CensorWord`, but you need to build the
  asterisk replacement string to match the rune count, not the byte count.
- `strings.Repeat("*", n)` repeats a string n times.
