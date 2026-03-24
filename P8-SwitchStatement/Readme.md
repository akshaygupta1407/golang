Switch Statement in Go

Go has a very clean and powerful `switch`. It is often easier to read than a long `if / else if / else` chain, especially when you are matching one value against many cases.

This folder (`P8-SwitchStatement`) shows the most common switch patterns in `main.go`, and this README expands on them with a deeper explanation of how switch works in Go.

From this example (`P8-SwitchStatement/main.go`)

Your program demonstrates:

- a basic value switch
- multiple values in one `case`
- a `switch` using `time.Now().Weekday()`
- a type switch using `interface{}`

Run it:

```bash
cd P8-SwitchStatement
go run .
```

------------------------------------------------------------

1) Basic value switch

The simplest form compares one expression against several possible values:

```go
i := 5

switch i {
case 1:
	fmt.Println("one")
case 2:
	fmt.Println("two")
case 3:
	fmt.Println("three")
case 4:
	fmt.Println("four")
case 5:
	fmt.Println("five")
default:
	fmt.Println("other")
}
```

How it works:

- `switch i` means "compare `i` against each case"
- the first matching `case` runs
- if nothing matches, `default` runs

This is useful when:

- mapping numbers to names
- mapping input commands to actions
- selecting logic for known constant values

------------------------------------------------------------

2) Multiple values in one case

Go lets you group multiple matching values into one `case`:

```go
switch time.Now().Weekday() {
case time.Saturday, time.Sunday:
	fmt.Println("It is weekend")
default:
	fmt.Println("It is weekday")
}
```

This is cleaner than writing separate cases when the action is the same.

Good use cases:

- weekend vs weekday
- vowels vs consonants
- multiple roles with the same permissions

------------------------------------------------------------

3) `default` branch

`default` is the fallback branch. It runs only when no `case` matches.

```go
switch value {
case 1:
	fmt.Println("matched 1")
default:
	fmt.Println("no match")
}
```

Notes:

- `default` is optional
- if you omit it and nothing matches, the switch simply does nothing

------------------------------------------------------------

4) No automatic fallthrough

One of the nicest things about Go switch statements is that cases do not fall through automatically.

In many other languages, if one case matches, execution may continue into the next case unless you write `break`. In Go, that does not happen by default.

Example:

```go
switch 2 {
case 2:
	fmt.Println("two")
case 3:
	fmt.Println("three")
}
```

Output:

```go
two
```

Only the matching case runs.

This makes switch blocks safer and easier to read.

------------------------------------------------------------

5) Manual `fallthrough`

If you really want execution to continue into the next case, Go provides `fallthrough`:

```go
switch 2 {
case 2:
	fmt.Println("two")
	fallthrough
case 3:
	fmt.Println("three")
}
```

Output:

```go
two
three
```

Important:

- `fallthrough` moves to the next case block without checking its condition
- it is used rarely in idiomatic Go

Most of the time, if logic is shared, it is better to group values in the same case.

------------------------------------------------------------

6) Expression-less switch

Go also supports `switch` without an expression:

```go
score := 82

switch {
case score >= 90:
	fmt.Println("A")
case score >= 75:
	fmt.Println("B")
default:
	fmt.Println("C")
}
```

This works like a cleaner `if / else if / else` chain.

It is especially useful for:

- numeric ranges
- validation rules
- feature flags
- branching on multiple boolean conditions

------------------------------------------------------------

7) Type switch

Your example includes a type switch:

```go
whoAmI := func(i interface{}) {
	switch t := i.(type) {
	case int:
		fmt.Println("It's an integer")
	case string:
		fmt.Println("It's a string")
	case bool:
		fmt.Println("It's a boolean")
	default:
		fmt.Println("It's another type", t)
	}
}
```

What this means:

- `i` is an `interface{}`
- `i.(type)` asks Go to inspect the dynamic type stored inside it
- each case matches a concrete type

This pattern is useful when:

- handling values of different types
- working with generic input like `any`
- decoding data from loose interfaces

Important note:

- type switches work only inside a `switch`
- they are different from normal value switches

------------------------------------------------------------

8) `interface{}` and `any`

In modern Go, `any` is just an alias for `interface{}`.

So this:

```go
func whoAmI(i interface{}) {}
```

can also be written as:

```go
func whoAmI(i any) {}
```

Both mean "this value can be of any type".

------------------------------------------------------------

9) When to use `switch` vs `if`

Prefer `switch` when:

- you compare one value against many options
- several conditions lead to clearly separated branches
- readability improves by grouping related cases

Prefer `if` when:

- there are only one or two branches
- the logic is simple and direct
- you are checking one special condition and returning early

------------------------------------------------------------

10) Best practices

- use `switch` instead of long repeated equality checks
- group related values in one case when the logic is identical
- avoid `fallthrough` unless you really need it
- use expression-less `switch` for clean range checks
- use type switches carefully when handling `interface{}` or `any`

------------------------------------------------------------

Next experiments

- Add a `switch` without an expression for age groups.
- Try replacing a long `if / else if` chain with `switch` and compare readability.
- Update the type switch to use `any` instead of `interface{}` and observe that the behavior stays the same.

