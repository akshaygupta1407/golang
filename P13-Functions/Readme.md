Functions in Go

Functions are one of the core building blocks of Go. They let you group logic, reuse behavior, and give names to operations that would otherwise be repeated inline.

This folder (`P13-Functions`) shows the basic function patterns in `main.go`, and this README adds a deeper explanation of how functions work in Go.

From this example (`P13-Functions/main.go`)

Your program demonstrates:

- defining a function with parameters
- returning a single value
- returning multiple values
- assigning function results to variables

Run it:

```bash
cd P13-Functions
go run .
```

------------------------------------------------------------

1) What a function is

A function is a named block of code that can accept inputs and optionally return outputs.

Example:

```go
func add(a int, b int) int {
	return a + b
}
```

This means:

- the function is named `add`
- it takes two `int` parameters
- it returns one `int`

Functions help you:

- avoid duplication
- make code easier to read
- separate responsibilities into smaller pieces

------------------------------------------------------------

2) Function signature

The part of the function that describes its inputs and outputs is called the signature.

For `add`:

```go
func add(a int, b int) int
```

Breakdown:

- `a int`, `b int` are parameters
- `int` at the end is the return type

In Go, types are written after the variable name, which makes function signatures very explicit.

------------------------------------------------------------

3) Calling a function

In `main`, you call the function like this:

```go
sum := add(3, 5)
fmt.Println(sum)
```

How it works:

- `3` is passed to `a`
- `5` is passed to `b`
- the return value is assigned to `sum`

This is the normal way to reuse logic in Go.

------------------------------------------------------------

4) Returning multiple values

Go functions can return more than one value.

Your example:

```go
func getLanguages() (string, string, string) {
	return "Go", "Python", "JavaScript"
}
```

This function returns three strings.

In `main`, you receive them like this:

```go
lang1, lang2, lang3 := getLanguages()
fmt.Println(lang1, lang2, lang3)
```

Multiple return values are one of Go’s most useful features. They are commonly used for:

- result + error
- value + status
- multiple related outputs

------------------------------------------------------------

5) Named return values

Your sample does not use named returns, but Go supports them:

```go
func split() (first string, second string) {
	first = "Go"
	second = "Lang"
	return
}
```

Named returns can make small functions easier to read, but they should be used carefully because they can also reduce clarity if overused.

------------------------------------------------------------

6) Parameter types

Go requires explicit parameter types.

Example:

```go
func add(a int, b int) int
```

You can also group parameters of the same type:

```go
func add(a, b int) int
```

This is a nice Go style when multiple arguments share the same type.

------------------------------------------------------------

7) Return statements

To send values back from a function, use `return`:

```go
return a + b
```

If a function declares return values, it must return compatible values along every path.

If a function does not return anything, its return type is omitted.

Example:

```go
func greet() {
	fmt.Println("hello")
}
```

------------------------------------------------------------

8) Functions as reusable units

Good functions usually do one clear thing.

`add` does one thing:

- it adds two integers

`getLanguages` does one thing:

- it returns a set of language names

That separation makes code easier to test, reuse, and modify.

------------------------------------------------------------

9) Common function patterns in Go

A) Utility functions

```go
func isEven(n int) bool {
	return n%2 == 0
}
```

B) Result plus error

```go
func parseAge(input string) (int, error) {
	// parse logic here
	return 0, nil
}
```

C) Multiple related outputs

```go
func dimensions() (int, int) {
	return 10, 20
}
```

D) Small helper functions

```go
func printHeader() {
	fmt.Println("=== Start ===")
}
```

------------------------------------------------------------

10) Best practices

- keep functions small and focused
- prefer clear parameter names
- use multiple returns for related outputs when it improves readability
- avoid overly clever one-liners if they hide intent
- use helper functions when a block of code deserves its own name

------------------------------------------------------------

11) What this example teaches

This chapter is a simple introduction, but it already shows a few core ideas:

- functions are declared with `func`
- parameters and return types are explicit
- functions can return one or many values
- function results can be unpacked into variables directly

------------------------------------------------------------

Next experiments

- Add a function that takes a `string` and returns its length.
- Create a function that returns `(value, ok)` like many Go APIs.
- Try a function with grouped parameters, like `func add(a, b int) int`.

