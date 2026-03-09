# Go (Golang) — Fundamentals Reference Guide

> **Goal:** This guide is structured for programmers with prior experience who are learning Go for the first time. The ultimate aim is to get comfortable enough to contribute to large open-source Go projects like [Kubernetes](https://github.com/kubernetes/kubernetes).

---

## Table of Contents

1. [Why Go?](#1-why-go)
2. [Packages and Modules](#2-packages-and-modules)
3. [Variables and Types](#3-variables-and-types)
4. [Functions](#4-functions)
5. [Control Flow](#5-control-flow)
6. [Arrays and Slices](#6-arrays-and-slices)
7. [Maps](#7-maps)
8. [Strings and Runes](#8-strings-and-runes)
9. [Structs and Interfaces](#9-structs-and-interfaces)
10. [Pointers](#10-pointers)
11. [Goroutines and Concurrency](#11-goroutines-and-concurrency)
12. [Channels](#12-channels)
13. [Generics](#13-generics)
14. [Exercises](#14-exercises)

---

## 1. Why Go?

Understanding *why* Go was designed the way it was helps you write idiomatic Go code.

| Feature | Description |
|---|---|
| **Statically Typed** | Types are checked at compile time — fewer runtime surprises. |
| **Strongly Typed** | You cannot mix types implicitly (e.g., cannot add a `string` and an `int`). |
| **Compiled to Machine Code** | Like C/C++/Rust — Go compiles directly to native machine code, not to a VM bytecode like Java or C#. |
| **Fast Compile Time** | Despite compiling to machine code, Go's compile times are extremely fast. |
| **Built-in Concurrency** | Goroutines and channels are first-class citizens of the language. |
| **Simplicity** | Go deliberately has a small feature set. If something feels overly complex in Go, you're probably doing it wrong. |

### Execution Speed Context

```
Compiled to VM bytecode  →  Java, C#        (slower startup, JIT warms up)
Compiled to Machine Code →  C, C++, Rust, Go (fast execution)
```

> **Interesting note:** Despite Go compiling to machine code (like C/C++), its execution speed is often *comparable* to JVM languages in practice — while its compile times are far faster than C/C++. This makes Go a sweet spot for infrastructure software like Kubernetes, Docker, and Terraform.

---

## 2. Packages and Modules

This is one of Go's key organizational concepts. Understand this early.

```
Package  = a collection of .go files in the same directory
Module   = a collection of packages (defined by go.mod)
```

### Rules You Must Know

- **One folder = one package.** All `.go` files in a folder must declare the same package name at the top.
- **Every import must be used.** Go will throw a *compile error* if you import a package and don't use it.
- **Every declared variable must be used.** Same rule — unused variables are compile errors.

```go
package main  // Every executable program must have a 'main' package

import (
    "fmt"
    // "math" ← This would cause a compile error if not used
)

func main() {
    fmt.Println("Hello, Go!")
}
```

> **Why these rules?** Go enforces clean code at the compiler level. No dead imports, no unused variables. This matters a lot in large codebases like Kubernetes.

---

## 3. Variables and Types

### Basic Types

| Go Type | Description | Equivalent in Other Languages |
|---|---|---|
| `int`, `int8`, `int16`, `int32`, `int64` | Signed integers | `int` in Java |
| `uint`, `uint8`, `uint16`, `uint32`, `uint64` | Unsigned integers | — |
| `float32`, `float64` | Floating point numbers | `float`, `double` |
| `string` | UTF-8 encoded text | `String` in Java |
| `bool` | `true` / `false` | `boolean` |
| `rune` | A single Unicode character (alias for `int32`) | `char` in Java |
| `byte` | Alias for `uint8` | `byte` |

### Declaring Variables — Three Ways

```go
// 1. Full declaration with explicit type
var myString string = "hello"

// 2. Type inference — Go figures out the type from the value
var myString = "hello"

// 3. Short declaration — most common inside functions
myString := "hello"
```

> **Rule of thumb:** Use `:=` inside functions. Use `var` at the package level.

### Constants

Constants must be assigned a value at declaration time — you cannot declare a constant without a value.

```go
const Pi = 3.14159
const MaxRetries int = 5

// This would be a compile error:
// const x int  ← no value assigned
```

---

## 4. Functions

### Basic Function

```go
// Go requires you to declare the return type explicitly
func add(a int, b int) int {
    return a + b
}

// Void functions (no return type needed)
func greet(name string) {
    fmt.Println("Hello,", name)
}
```

### Multiple Return Values

This is one of Go's most distinctive features — and it's used *everywhere* in real Go code.

```go
func divide(a int, b int) (int, int) {
    return a / b, a % b
}

quotient, remainder := divide(10, 3)
fmt.Println(quotient, remainder) // 3 1
```

### The Error Pattern — Learn This Well

A very common Go design pattern is to return an `error` as the last return value. You will see this constantly in Kubernetes source code.

```go
import "errors"

func intDivision(numerator int, denominator int) (int, int, error) {
    var err error

    if denominator == 0 {
        err = errors.New("cannot divide by zero")
        return 0, 0, err
    }

    result := numerator / denominator
    remainder := numerator % denominator
    return result, remainder, err
}

func main() {
    result, remainder, err := intDivision(10, 0)
    if err != nil {
        fmt.Printf("Error: %v\n", err.Error())
        return
    }
    fmt.Printf("Result: %v, Remainder: %v\n", result, remainder)
}
```

> **Key insight:** In Go, errors are values — not exceptions. You check `if err != nil` after almost every function call. This explicit error handling is a deliberate Go philosophy.

---

## 5. Control Flow

### If / Else

```go
x := 10
if x > 5 {
    fmt.Println("Greater")
} else if x == 5 {
    fmt.Println("Equal")
} else {
    fmt.Println("Less")
}
```

### Switch — No Explicit `break` Needed

Unlike Java/C, Go's `switch` cases break automatically. You don't need to write `break`.

```go
day := "Monday"
switch day {
case "Monday":
    fmt.Println("Start of the week")
case "Friday":
    fmt.Println("End of the week")
default:
    fmt.Println("Midweek")
}
```

> Go also supports **short-circuit evaluation** in boolean expressions, just like Java — `&&` stops if the left side is `false`, `||` stops if the left side is `true`.

### For Loop — Go's Only Loop

Go has **no `while` loop**. The `for` keyword covers all looping patterns.

```go
// Traditional for loop
for i := 0; i < 10; i++ {
    fmt.Println(i)
}

// Acts like a while loop
i := 0
for i < 10 {
    fmt.Println(i)
    i++
}

// Range-based loop (like for-each)
nums := []int{1, 2, 3, 4, 5}
for index, value := range nums {
    fmt.Printf("Index: %v, Value: %v\n", index, value)
}

// Ignore index with blank identifier
for _, value := range nums {
    fmt.Println(value)
}
```

---

## 6. Arrays and Slices

### Arrays — Fixed Size

Arrays in Go are **contiguous in memory** and have a **fixed size**. The size is part of the type — `[3]int` and `[5]int` are different types.

```go
var arr [5]int = [5]int{1, 2, 3, 4, 5}
fmt.Println(arr[0]) // 1

// Compiler can infer size with ...
arr2 := [...]string{"a", "b", "c"}
```

> **Why contiguous memory matters:** Because elements are next to each other in memory, the compiler doesn't need to store the address of each element. It can calculate any element's address as `base_address + (index * element_size)`. This makes array access O(1) and very cache-friendly.

### Slices — Dynamic, Flexible Arrays

Slices are the workhorse of Go. Think of them as a **view into an underlying array** with extra capabilities — similar to Python lists, but homogeneous (all elements must be the same type).

```go
// Creating a slice
var mySlice []int = []int{1, 2, 3}

// Using make (length=3, capacity=5)
mySlice2 := make([]int, 3, 5)

fmt.Println(len(mySlice))  // length: number of elements currently in slice
fmt.Println(cap(mySlice))  // capacity: how many elements it can hold before resizing
```

### Length vs Capacity

```
Underlying Array:  [1, 2, 3, _, _]
                    ^--------^        <- length = 3
                    ^-----------^     <- capacity = 5
```

### Appending to Slices

```go
mySlice := []int{1, 2, 3}
mySlice = append(mySlice, 4, 5)
fmt.Println(mySlice) // [1 2 3 4 5]
```

> ⚠️ **Important:** Slices are **reference types**. When you pass a slice to a function, you're passing a reference to the underlying array. Modifying the slice inside the function **modifies the original data**. More on this in the Pointers section.

---

## 7. Maps

Maps are Go's built-in key-value data structure (like `HashMap` in Java or `dict` in Python), but **homogeneous** — all keys must be the same type, all values must be the same type.

```go
// Declaring and initializing a map
var ages = map[string]uint8{
    "Adam":  23,
    "Sarah": 45,
}

// Accessing a value
fmt.Println(ages["Adam"]) // 23
```

### Safe Key Lookup — The Two-Value Return

If you access a key that doesn't exist, Go returns the **zero value** for the type (e.g., `0` for `int`). To check if a key actually exists, use the two-value form:

```go
age, ok := ages["Jason"]
if ok {
    fmt.Printf("Jason's age is %v\n", age)
} else {
    fmt.Println("Jason not found")
}
```

### Deleting from a Map

```go
delete(ages, "Adam")
```

### Iterating Over a Map

```go
for name, age := range ages {
    fmt.Printf("Name: %v, Age: %v\n", name, age)
}
```

> **Note:** Map iteration order is **not guaranteed** in Go. Do not rely on insertion order.

---

## 8. Strings and Runes

This section has some depth — understanding it will prevent subtle bugs.

### Strings are UTF-8 Byte Arrays

In Go, a `string` is stored as an **underlying array of bytes** encoded in UTF-8. This has an important consequence: **indexing a string gives you a byte, not a character.**

```go
s := "résumé"

// This gives you a byte (uint8), NOT a character
fmt.Println(s[1]) // 195 — just the first byte of 'é', not 'é' itself
```

### Why indexing can break on non-ASCII characters

`é` in UTF-8 is encoded as **two bytes**: `11000011 10101001` (i.e., bytes 195 and 169). So when you index position 1, you only get `195` — half of the character.

```
s = "résumé"
Bytes: [r][é byte1][é byte2][s][u][m][é byte1][é byte2]
Index:  0     1        2     3  4  5     6        7
```

Index 2 gets skipped in a `for range` loop because the loop correctly handles multi-byte characters.

### Iterating Correctly with `range`

```go
s := "résumé"

// CORRECT: range iterates by rune (Unicode code point), not byte
for index, char := range s {
    fmt.Printf("Index: %v, Char: %c\n", index, char)
}
```

### Converting to Rune Slice

The easiest way to work with individual characters is to convert to `[]rune`:

```go
s := "résumé"
runes := []rune(s)
fmt.Println(string(runes[1])) // é — correct!
```

> **What is a rune?** A `rune` is just an alias for `int32`. It represents a Unicode code point. Think of it as Go's version of `char` in Java — but capable of representing *any* Unicode character, not just ASCII.

### String Immutability and `strings.Builder`

Strings in Go are **immutable**, just like in Java and Python. Each concatenation with `+` creates a new string.

```go
// Inefficient — creates many intermediate strings
result := ""
for i := 0; i < 1000; i++ {
    result += "a"
}

// Efficient — use strings.Builder
import "strings"

var sb strings.Builder
for i := 0; i < 1000; i++ {
    sb.WriteString("a")
}
result := sb.String()
```

> `strings.Builder` internally uses a byte buffer and only converts to a string at the end — much more efficient for repeated concatenation.

---

## 9. Structs and Interfaces

### Structs — Defining Your Own Types

A `struct` lets you define a custom data type by grouping fields together. Similar to classes in Java/Python, but Go structs have no inheritance.

```go
type gasEngine struct {
    mpg     uint8
    gallons uint8
}

func main() {
    // Initialize with field names (recommended)
    myEngine := gasEngine{mpg: 25, gallons: 15}
    fmt.Println(myEngine.mpg, myEngine.gallons) // 25 15
}
```

### Nested Structs

```go
type owner struct {
    name string
}

type gasEngine struct {
    mpg     uint8
    gallons uint8
    owner   owner // nested struct
}

myEngine := gasEngine{
    mpg:     25,
    gallons: 15,
    owner:   owner{name: "Alice"},
}
fmt.Println(myEngine.owner.name) // Alice
```

### Anonymous Structs

One-off structs that you don't need to reuse elsewhere:

```go
myEngine := struct {
    mpg     uint8
    gallons uint8
}{25, 15}
```

### Methods on Structs

You can attach functions to structs. This is Go's answer to class methods.

```go
type gasEngine struct {
    mpg     uint8
    gallons uint8
}

// Method with a receiver — (e gasEngine) means this method belongs to gasEngine
func (e gasEngine) milesLeft() uint8 {
    return e.gallons * e.mpg
}

func main() {
    myEngine := gasEngine{mpg: 25, gallons: 15}
    fmt.Printf("Miles left: %v\n", myEngine.milesLeft()) // 375
}
```

> **Your question answered:** `myEngine.milesLeft()` works because `milesLeft` has a *receiver* of type `gasEngine`. This is Go's way of associating a function with a type. It is not inside the struct definition, but it "belongs" to it via the receiver.

### Interfaces

An interface defines a set of method signatures. Any type that implements all those methods *automatically* satisfies the interface — no explicit `implements` keyword needed (unlike Java).

```go
type engine interface {
    milesLeft() uint8
}

type gasEngine struct {
    mpg     uint8
    gallons uint8
}

type electricEngine struct {
    kwh   float32
    mpkwh float32
}

func (e gasEngine) milesLeft() uint8 {
    return uint8(e.gallons * e.mpg)
}

func (e electricEngine) milesLeft() uint8 {
    return uint8(e.kwh * e.mpkwh)
}

// This function accepts ANY type that implements the engine interface
func canMakeIt(e engine, miles uint8) {
    if miles <= e.milesLeft() {
        fmt.Println("You can make it!")
    } else {
        fmt.Println("Not enough fuel/charge.")
    }
}

func main() {
    gas := gasEngine{mpg: 25, gallons: 10}
    electric := electricEngine{kwh: 57.5, mpkwh: 4.17}

    canMakeIt(gas, 200)
    canMakeIt(electric, 200)
}
```

> **This is huge.** In Kubernetes, interfaces are used *extensively*. For example, a storage backend might be swapped between cloud providers — each implementing the same interface. Your code using `canMakeIt` doesn't care whether it receives a `gasEngine` or `electricEngine`, as long as it has a `milesLeft()` method.

---

## 10. Pointers

### What is a Pointer?

A pointer is a variable that stores the **memory address** of another variable.

```go
var p *int32       // p is a pointer to an int32 (currently nil)
var i int32 = 42

p = &i             // & gives us the memory address of i
fmt.Println(*p)    // * dereferences the pointer — gives us the value: 42

*p = 100           // Modifying via pointer also changes i
fmt.Println(i)     // 100
```

### Pointer Operators Summary

| Operator | Meaning |
|---|---|
| `*T` | Type: "pointer to T" |
| `&x` | Address of variable `x` |
| `*p` | Dereference: value at the address stored in `p` |
| `new(T)` | Allocates memory for a T, returns a pointer to it |

### Why Use Pointers?

**Without pointer** — function gets a copy:
```go
func double(x int) {
    x = x * 2 // only modifies the local copy
}

n := 5
double(n)
fmt.Println(n) // still 5
```

**With pointer** — function modifies the original:
```go
func double(x *int) {
    *x = *x * 2 // modifies the original
}

n := 5
double(&n)
fmt.Println(n) // 10
```

**Passing large arrays efficiently:**
```go
// WITHOUT pointer — copies entire array (expensive)
func square(arr [5]float64) [5]float64 { ... }

// WITH pointer — passes just the address (cheap)
func square(arr *[5]float64) {
    for i := range arr {
        arr[i] = arr[i] * arr[i]  // modifies original
    }
}

var thing1 = [5]float64{1, 2, 3, 4, 5}
square(&thing1)
fmt.Println(thing1) // [1 4 9 16 25]
```

### Slices are Already Reference Types

Slices internally contain a pointer to their underlying array. This means **passing a slice to a function already gives the function access to the original data** — no need for explicit pointers.

```go
func zeroOut(s []int) {
    s[0] = 0 // modifies the original slice's data
}

nums := []int{1, 2, 3}
zeroOut(nums)
fmt.Println(nums) // [0 2 3]
```

---

## 11. Goroutines and Concurrency

### Concurrency vs Parallelism

> "Concurrency is about *dealing with* lots of things at once. Parallelism is about *doing* lots of things at once." — Rob Pike (Go co-creator)

- **Concurrent:** Tasks are in progress simultaneously — they may take turns on a single core.
- **Parallel:** Tasks are literally executing at the same moment on multiple CPU cores.
- All parallel execution is concurrent, but not all concurrency is parallel.

### Launching a Goroutine

Add the `go` keyword before a function call to run it concurrently:

```go
go myFunction()
```

### WaitGroup — Waiting for Goroutines to Finish

Without a `WaitGroup`, your `main` function may exit before goroutines finish.

```go
import (
    "fmt"
    "math/rand"
    "sync"
    "time"
)

var wg = sync.WaitGroup{}
var dbData = []string{"id1", "id2", "id3", "id4", "id5"}

func main() {
    t0 := time.Now()

    for i := 0; i < len(dbData); i++ {
        wg.Add(1)       // increment counter before launching
        go dbCall(i)
    }

    wg.Wait()           // block until counter reaches 0
    fmt.Printf("Total execution time: %v\n", time.Since(t0))
}

func dbCall(i int) {
    defer wg.Done()     // decrement counter when function returns

    delay := rand.Float32() * 2000
    time.Sleep(time.Duration(delay) * time.Millisecond)
    fmt.Println("Result:", dbData[i])
}
```

> **`defer`** schedules a function call to run when the surrounding function returns — great for cleanup and `wg.Done()` calls.

### Mutex — Preventing Race Conditions

When multiple goroutines write to a shared variable, you get a **race condition** — unpredictable, hard-to-debug bugs.

```go
var m = sync.Mutex{}
var results = []string{}

func save(result string) {
    m.Lock()                          // only one goroutine can proceed past here
    results = append(results, result)
    m.Unlock()                        // release the lock
}
```

### RWMutex — Allowing Concurrent Reads

`sync.Mutex` blocks *everyone*. `sync.RWMutex` allows **multiple simultaneous readers** but only **one writer at a time** — better performance when reads are frequent.

```go
var m = sync.RWMutex{}

func save(result string) {
    m.Lock()
    results = append(results, result)
    m.Unlock()
}

func log() {
    m.RLock()                         // multiple goroutines can RLock simultaneously
    fmt.Printf("Current results: %v\n", results)
    m.RUnlock()
}
```

> **Where to put Lock/Unlock matters a lot.** Locking too broadly kills concurrency gains. Locking too narrowly causes race conditions. In Kubernetes, this is a common source of subtle bugs.

### Performance Note

Goroutine performance gains depend on the workload:
- **I/O-bound tasks** (network calls, DB queries): Massive gains — goroutines shine here.
- **CPU-bound tasks**: Gains are proportional to the number of CPU cores available.
- **Lightweight tasks**: Overhead of goroutine management may outweigh the gains.

---

## 12. Channels

Channels are Go's mechanism for **communication between goroutines**. They are:
1. Type-safe containers for passing data
2. Thread-safe (no manual locking needed)
3. Able to block goroutines until data is available

```go
// Create a channel that carries int values
var c = make(chan int)
```

### Sending and Receiving

```go
c <- value   // Send value INTO the channel
value := <-c // Receive value OUT OF the channel
```

### Example — Producer/Consumer

```go
func main() {
    c := make(chan int)
    go produce(c)

    for val := range c {  // receives until channel is closed
        fmt.Println(val)
    }
}

func produce(c chan int) {
    for i := 0; i < 5; i++ {
        c <- i
    }
    close(c)  // IMPORTANT: always close the channel when done sending
}
```

> ⚠️ **Forgetting to close a channel causes a deadlock.** The receiver will wait forever for more values that never come.

### Real-World Pattern — First Result Wins

This pattern is used when you query multiple sources and want the fastest response:

```go
const MAX_CHICKEN_PRICE float32 = 5.0

func main() {
    chickenChannel := make(chan string)
    websites := []string{"walmart.com", "costco.com", "wholefoods.com"}

    for _, site := range websites {
        go checkPrice(site, chickenChannel)
    }

    // Blocks until the FIRST goroutine sends a result
    fmt.Printf("Best deal found at: %s\n", <-chickenChannel)
}

func checkPrice(website string, ch chan string) {
    for {
        time.Sleep(time.Second * 1)
        price := rand.Float32() * 20
        if price <= MAX_CHICKEN_PRICE {
            ch <- website
            break
        }
    }
}
```

> This pattern is commonly used in distributed systems — fire off requests to multiple replicas or APIs and use whichever responds first.

---

## 13. Generics

Introduced in Go 1.18, generics allow you to write functions and types that work with multiple types without sacrificing type safety.

### Generic Functions

```go
// T can be int, float32, or float64
func sumSlice[T int | float32 | float64](slice []T) T {
    var sum T
    for _, v := range slice {
        sum += v
    }
    return sum
}

func main() {
    ints := []int{1, 2, 3}
    floats := []float32{1.1, 2.2, 3.3}

    fmt.Println(sumSlice[int](ints))         // 6
    fmt.Println(sumSlice[float32](floats))   // 6.6
}
```

> **Why not use `any`?** The `any` type (alias for `interface{}`) disables type checking — you lose safety and can't use operators like `+`. Generics give you flexibility *with* type safety.

### Generic Structs

```go
type gasEngine struct {
    gallons float32
    mpg     float32
}

type electricEngine struct {
    kwh   float32
    mpkwh float32
}

// car can hold either a gasEngine or electricEngine
type car[T gasEngine | electricEngine] struct {
    carMake  string
    carModel string
    engine   T
}

func main() {
    gasCar := car[gasEngine]{
        carMake:  "Honda",
        carModel: "Civic",
        engine:   gasEngine{gallons: 12.4, mpg: 40},
    }

    electricCar := car[electricEngine]{
        carMake:  "Tesla",
        carModel: "Model 3",
        engine:   electricEngine{kwh: 57.5, mpkwh: 4.17},
    }

    fmt.Println(gasCar)
    fmt.Println(electricCar)
}
```

---

## 14. Exercises

These exercises are ordered by difficulty. Complete them alongside these notes. Each maps to a concept above.

### Beginner

1. **Variables & Types:** Declare variables of at least 6 different types. Print each one with its type using `fmt.Printf("%T", variable)`.

2. **Error Handling:** Write a function `safeSqrt(x float64) (float64, error)` that returns an error if `x` is negative.

3. **Slices:** Write a function that takes a slice of integers and returns a new slice with only the even numbers.

4. **Maps:** Build a word frequency counter — given a sentence string, return a `map[string]int` with the count of each word.

5. **Structs:** Define a `Rectangle` struct with `width` and `height` fields. Add methods `Area()` and `Perimeter()`.

### Intermediate

6. **Interfaces:** Define a `Shape` interface with `Area() float64` and `Perimeter() float64`. Implement it for `Rectangle`, `Circle`, and `Triangle`. Write a function that takes a `[]Shape` and prints the total area.

7. **Pointers:** Write two versions of a function that doubles every element in an integer array — one that copies, one that modifies in place using a pointer. Benchmark or observe the difference.

8. **Goroutines:** Simulate fetching data from 5 "databases" with random delays using goroutines. Print results as they arrive and measure total vs sequential execution time.

9. **Channels:** Implement a simple pipeline: one goroutine generates numbers 1–20, another filters out odd numbers, a third prints the results. Connect them with channels.

10. **Strings/Runes:** Write a function that reverses a string correctly, including strings with multi-byte Unicode characters like `"résumé"`. (Hint: convert to `[]rune` first.)

### Advanced (Kubernetes-Oriented)

11. **Interfaces + Goroutines:** Build a mock health checker. Define a `Service` interface with `Name() string` and `IsHealthy() bool`. Implement 3 fake services. Concurrently check all of them and print a health report.

12. **RWMutex:** Build a simple in-memory key-value store (like a mini `etcd`) with concurrent-safe `Get`, `Set`, and `Delete` operations. Use `sync.RWMutex`.

13. **Generics:** Write a generic `Filter[T]` function that takes a slice and a predicate function `func(T) bool` and returns a filtered slice.

14. **Error Wrapping:** Research `fmt.Errorf` with `%w` and `errors.Is` / `errors.As`. Rewrite Exercise 2 to use wrapped errors and practice unwrapping them.

15. **Full Mini-Project:** Build a concurrent URL health checker. Given a list of URLs, check each one concurrently using goroutines, use a channel to collect results, and print a table showing URL, status code, and response time. This mimics patterns used in Kubernetes probes.

---

## Quick Reference Cheat Sheet

```go
// Variable declaration
x := 42
var y string = "hello"

// Function with error return
func fn() (int, error) { return 0, nil }

// Goroutine + WaitGroup
var wg sync.WaitGroup
wg.Add(1)
go func() { defer wg.Done(); doWork() }()
wg.Wait()

// Channel
ch := make(chan int)
go func() { ch <- 42; close(ch) }()
val := <-ch

// Mutex
var mu sync.Mutex
mu.Lock()
// critical section
mu.Unlock()

// Interface
type Doer interface { Do() error }

// Generic function
func Map[T, U any](s []T, f func(T) U) []U { ... }
```

---

*Last updated while learning Go with the goal of contributing to Kubernetes. Keep going — the best way to learn is to read real Go source code at [github.com/kubernetes/kubernetes](https://github.com/kubernetes/kubernetes).*