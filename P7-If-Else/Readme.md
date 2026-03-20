If Else in Go

Go keeps conditional logic simple and explicit. There is no `switch` requirement for basic branching, and there is no ternary operator. In practice, `if`, `else if`, and `else` cover most everyday decisions cleanly.

This folder (`P7-If-Else`) shows the core patterns in `main.go`, and this README adds a deeper walkthrough with a few useful Go-specific notes.

From this example (`P7-If-Else/main.go`)

Your program demonstrates:

- a basic age-based decision tree
- boolean conditions with `||` and `&&`
- declaring a variable inside an `if` statement
- using `if` as the replacement for a ternary expression

Run it:

```bash
cd P7-If-Else
go run .
```

------------------------------------------------------------

1) Basic `if`, `else if`, `else`

The most common shape is a decision ladder:

```go
age := 16
if age >= 18 {
	fmt.Println("person is an adult")
} else if age >= 12 {
	fmt.Println("person is a teenager")
} else {
	fmt.Println("person is a kid")
}
```

How it works:

- Go evaluates the conditions from top to bottom
- the first true branch runs
- the remaining branches are skipped

This is a natural fit for:

- age bands
- score ranges
- validation checks
- state classification

------------------------------------------------------------

2) Boolean logic in conditions

Your example also shows logical operators:

```go
if role == "admin" || hasPermission {
	fmt.Println("yes user is have access")
}

if role == "admin" && hasPermission {
	fmt.Println("yes user is have access")
}
```

`||` means OR:

- the condition is true if either side is true

`&&` means AND:

- the condition is true only if both sides are true

Useful reminder:

- `&&` has higher precedence than `||`
- when conditions get complex, add parentheses for readability

Example:

```go
if (role == "admin" || role == "manager") && hasPermission {
	fmt.Println("allowed")
}
```

------------------------------------------------------------

3) Short variable declaration inside `if`

Go lets you declare a variable directly in the `if` header:

```go
if age := 15; age >= 18 {
	fmt.Println("person is an adult. age:", age)
} else if age >= 12 {
	fmt.Println("person is a teenager. age:", age)
}
```

This is very useful when:

- a value is only needed for one decision
- you want to keep the variable scoped tightly
- you want to avoid cluttering the outer function scope

Important scope rule:

- `age` exists only inside the `if / else if / else` chain
- it is not available after the statement ends

This pattern is common with:

- parsed values
- lookups from maps
- error checks from function calls

Example:

```go
if n, err := strconv.Atoi(input); err == nil {
	fmt.Println("parsed:", n)
}
```

------------------------------------------------------------

4) Scope and readability

Go prefers explicit scopes. That means you often choose between:

- a variable declared before the `if`
- a short variable declared inside the `if`

Use the narrower scope when possible. It helps readers see that the variable is only relevant to one branch of logic.

Good rule of thumb:

- use a short `if` initializer for one-off checks
- use a broader variable when multiple statements need it later

------------------------------------------------------------

5) No ternary operator

Go does not have a C-style ternary operator like `condition ? a : b`.

Instead, write a normal `if / else`:

```go
result := ""
if age >= 18 {
	result = "adult"
} else {
	result = "minor"
}
```

Why this is considered a feature:

- it keeps control flow readable
- it avoids dense one-line expressions
- it makes side effects easier to see

If the expression is simple, you can keep the `if` short:

```go
status := "minor"
if age >= 18 {
	status = "adult"
}
```

------------------------------------------------------------

6) Common patterns in real code

A) Guard clauses

Use early returns to reduce nesting:

```go
if err != nil {
	return err
}
```

B) Validation

```go
if username == "" {
	return fmt.Errorf("username is required")
}
```

C) Feature access

```go
if role == "admin" && hasPermission {
	fmt.Println("access granted")
} else {
	fmt.Println("access denied")
}
```

D) Multi-way classification

```go
switch {
case score >= 90:
	fmt.Println("A")
case score >= 75:
	fmt.Println("B")
default:
	fmt.Println("C")
}
```

Note:

- for many range-based comparisons, `switch` can be cleaner than a long `if / else if` chain
- for simple binary decisions, `if` is usually best

------------------------------------------------------------

7) Best practices

- keep conditions readable
- avoid deeply nested `if` blocks when an early return works better
- use parentheses when boolean logic becomes hard to scan
- prefer clear branch names over clever one-liners
- use `if` initializers when the variable should not leak outside the decision

------------------------------------------------------------

Next experiments

- Add a `switch` example for the same age classification and compare readability.
- Add an input validation example that uses `if value := ...; value != ...`.
- Try converting one of the access checks into a guard-clause style function.

