# 04 — Maps

## Concept

Maps are Go's built-in hash table / dictionary. They store key-value pairs where
all keys must be the same type and all values must be the same type.

```go
m := map[string]int{"alice": 90, "bob": 85}
m["carol"] = 92           // add or update
delete(m, "bob")           // remove
val, ok := m["alice"]      // safe lookup: ok is true if key exists
```

Important behaviors:
- Accessing a missing key returns the **zero value** of the value type (not an error).
- Use the two-value form `val, ok := m[key]` to distinguish "key missing" from
  "key present with zero value".
- Map iteration order is **random** — never rely on insertion order.
- Maps are reference types — passing one to a function shares the underlying data.

## Exercise

Open `exercise.go`. You'll implement three functions:

1. **`WordFrequency(sentence string)`** — Count how many times each word appears
   in a sentence. This is a classic map exercise that shows up in real-world text
   processing.

2. **`MergeMaps(a, b map[string]int)`** — Merge two maps. If both maps have the
   same key, sum their values. Returns a new map without modifying the originals.

3. **`InvertMap(m map[string]string)`** — Swap keys and values. If two keys have
   the same value, only one will survive (that's expected).

Run with: `go run exercise.go`

## Hint

- `strings.Fields(s)` splits a string by whitespace — cleaner than `strings.Split`.
- `strings.ToLower(s)` normalizes case for fair word counting.
- When merging, iterate over both maps and accumulate into a new one.
- Think about what happens when you invert a map with duplicate values.
