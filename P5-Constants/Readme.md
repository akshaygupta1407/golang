Constants in Go

Constants are values that do not change after they are declared.
In Go, constants are created using the `const` keyword.

Example:

```go
const age = 30
```

Here, `age` will always stay `30`.

Why use constants

- Use constants for fixed values such as app names, ports, limits, or messages.
- They make code easier to read.
- They prevent accidental reassignment.

Constants cannot be changed

```go
const name string = "Hello"
```

If you try to assign a new value later:

```go
name = "Hi"
```

Go will give an error because constants are immutable.

Constants at package level

```go
const age = 30
```

This constant is declared outside `main()`, so it is available at package level.

One important difference:

- `const` can be used at package level
- `:=` cannot be used at package level

That is why this is invalid:

```go
// name1 := "fjf"
```

Constants inside a function

You can also declare constants inside functions:

```go
func main() {
	const name string = "Hello"
	fmt.Println(name)
}
```

Grouped constants

Go allows grouping multiple constants together:

```go
const (
	host = "localhost"
	port = 5000
)
```

This is useful when related constant values belong together.

From this example

Your program demonstrates:

- a package-level constant: `age`
- a typed constant: `const name string = "Hello"`
- grouped constants: `host` and `port`
- printing constant values with `fmt.Println`

Output:

```go
Hello
localhost 5000
```
