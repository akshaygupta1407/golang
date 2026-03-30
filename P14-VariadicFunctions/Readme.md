Variadic Functions in Go

A variadic function is a function that can accept a variable number of arguments. In Go, you define that with `...` in the parameter list. This is especially useful when you want a function to work naturally with zero, one, or many values.

This folder (`P14-VariadicFunctions`) shows the core variadic patterns in `main.go`, and this README adds the deeper theory behind how variadic parameters work in Go.

From this example (`P14-VariadicFunctions/main.go`)

Your program demonstrates:

- defining a variadic function
- passing multiple arguments directly
- passing a slice with `...`
- summing a flexible list of integers

Run it:

```bash
cd P14-VariadicFunctions
go run .
```

------------------------------------------------------------

1) What variadic means

Variadic means "variable number of arguments".

Your function:

```go
func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}
```

This means:

- `sum` can accept any number of `int` arguments
- inside the function, `nums` behaves like a slice of `int`

You can call it with:

```go
sum(1, 2, 3, 4, 5)
sum()
sum(10)
```

------------------------------------------------------------

2) Variadic parameters are slices inside the function

Inside the function body, the variadic parameter is treated like a slice.

So in:

```go
func sum(nums ...int) int
```

`nums` behaves like:

```go
[]int
```

That is why this works:

```go
for _, num := range nums {
	total += num
}
```

This is an important Go detail:

- variadic input is flexible at the call site
- but inside the function, you still work with slice-like behavior

------------------------------------------------------------

3) Calling a variadic function

Your sample calls `sum` directly with separate values:

```go
fmt.Println(sum(1, 2, 3, 4, 5))
```

This is the most common way to use a variadic function.

The compiler packages the arguments into the slice-like parameter automatically.

------------------------------------------------------------

4) Expanding a slice with `...`

Your example also shows:

```go
nums := []int{1, 2, 3, 4, 5}
fmt.Println(sum(nums...))
```

The `...` after a slice means "expand this slice into individual arguments".

This is useful when:

- you already have a slice
- you want to pass its contents to a variadic function
- you do not want to loop manually just to forward values

Without `...`, `sum(nums)` would not compile because `sum` expects `int` arguments, not a single `[]int`.

------------------------------------------------------------

5) Why variadic functions are useful

Variadic functions are great for APIs where the number of values is naturally flexible.

Common examples:

- logging functions
- formatting functions
- summing or aggregating data
- building helper utilities

For example, `fmt.Println` itself is variadic, which is why this works:

```go
fmt.Println(1, 2, 3, 4, 5, "hello")
```

------------------------------------------------------------

6) Variadic vs slice parameter

A variadic function and a slice parameter are related, but not identical.

Variadic form:

```go
func sum(nums ...int)
```

Slice form:

```go
func sum(nums []int)
```

Differences:

- variadic calls are nicer when passing separate values
- slice parameters are clearer when the caller already has a slice

Use variadic parameters when the function is naturally "any number of values".
Use slice parameters when the function conceptually works on one collection.

------------------------------------------------------------

7) The single variadic parameter rule

In Go, a function can have only one variadic parameter, and it must be the last parameter.

Example:

```go
func example(prefix string, nums ...int) {}
```

This is valid.

But this is not:

```go
func example(nums ...int, prefix string) {}
```

That rule keeps function calls unambiguous.

------------------------------------------------------------

8) Common patterns

A) Summation

```go
func sum(nums ...int) int
```

B) Logging

```go
func log(message string, args ...any)
```

C) Formatting

Many formatting helpers accept a variable list of arguments.

D) Forwarding arguments

You can receive a slice and pass it onward using `...`.

------------------------------------------------------------

9) Gotchas

- inside the function, variadic input is still a slice
- you can pass zero arguments, so handle empty input if needed
- the variadic parameter must be the last parameter
- use `slice...` only when the target function is variadic

------------------------------------------------------------

10) Best practices

- use variadic functions when the call site should stay simple
- prefer slices when the data is already grouped as a collection
- keep the function focused on one kind of aggregation or action
- handle empty input safely
- use `...any` for flexible logging or formatting helpers when appropriate

------------------------------------------------------------

11) What this example teaches

This chapter shows three practical ideas:

- how to declare a variadic function with `...`
- how to call it with multiple values
- how to forward an existing slice using `slice...`

That makes variadic functions one of the most useful small features in Go.

------------------------------------------------------------

Next experiments

- Add a variadic function that returns the average of numbers.
- Try a variadic function with a prefix parameter, like `func join(prefix string, parts ...string)`.
- Compare `func sum(nums ...int)` with `func sum(nums []int)` in a small example.

