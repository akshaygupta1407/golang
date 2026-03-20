Loops in Go (Deep Dive)

In Go, there is only one looping keyword: `for`.
Go does not have `while` or `do-while`. Instead, you express all looping patterns using `for`.

This folder (`P6-Loops`) shows a few basic patterns in `main.go`, and this README expands with more patterns, pitfalls, and best practices.

From this example (`P6-Loops/main.go`)

Your program demonstrates:

- a “while-style” loop using `for condition { ... }`
- a classic 3-part loop: `for init; condition; post { ... }`
- stopping a loop early with `break`
- `range` over an integer (Go 1.22+): `for i := range 3 { ... }`

Run it:

```bash
cd P6-Loops
go run .
```

------------------------------------------------------------

1) The 4 main loop shapes in Go

A) Classic 3-clause loop (like C/Java)

```go
for i := 0; i < 3; i++ {
	fmt.Println(i)
}
```

- `i := 0` runs once (init)
- `i < 3` is checked before each iteration (condition)
- `i++` runs after each iteration (post)

B) While-style loop (condition-only)

```go
i := 1
for i <= 3 {
	fmt.Println(i)
	i++
}
```

This is exactly what your `main.go` is doing in the “while loop” section.

C) Infinite loop

```go
for {
	// runs forever unless you break/return/panic
}
```

Typical uses:

- servers that “run forever”
- retry loops with backoff
- event loops with `select`

D) Range loop (iterate over a collection / string / channel / integer)

```go
nums := []int{10, 20, 30}
for idx, val := range nums {
	fmt.Println(idx, val)
}
```

You can omit variables you don’t need:

```go
for _, val := range nums { // ignore index
	fmt.Println(val)
}
```

------------------------------------------------------------

2) Loop control: `break`, `continue`, labels

break

- exits the nearest `for` loop

```go
for i := 0; i < 5; i++ {
	if i == 2 {
		break
	}
}
```

continue

- skips to the next iteration

```go
for i := 0; i < 5; i++ {
	if i%2 == 0 {
		continue
	}
	fmt.Println("odd:", i)
}
```

Labels (break/continue outer loops)

Labels help when you have nested loops and want to stop/continue an outer loop:

```go
outer:
for i := 0; i < 3; i++ {
	for j := 0; j < 3; j++ {
		if i == 1 && j == 1 {
			break outer
		}
	}
}
```

Notes:

- `break outer` breaks the loop labeled `outer`
- `continue outer` continues the loop labeled `outer`

------------------------------------------------------------

3) `range` deep dive

A) `range` over arrays/slices

```go
nums := []int{10, 20, 30}
for i, v := range nums {
	fmt.Println(i, v)
}
```

Important behaviors:

- `i` is the index (0..len-1)
- `v` is a copy of `nums[i]` (for slices of non-pointer values)

If you want to modify the slice elements, write through the index:

```go
for i := range nums {
	nums[i] *= 2
}
```

B) `range` over strings (bytes vs runes)

In Go, strings are bytes, but `range` decodes UTF-8 into runes:

```go
s := "hé"
for idx, r := range s {
	fmt.Println(idx, r) // idx is byte index; r is a rune
}
```

If you want raw bytes:

```go
for i := 0; i < len(s); i++ {
	fmt.Println(i, s[i]) // byte value
}
```

C) `range` over maps

```go
m := map[string]int{"a": 1, "b": 2}
for k, v := range m {
	fmt.Println(k, v)
}
```

Key notes:

- map iteration order is not guaranteed; do not rely on it
- if you need stable order, extract keys, sort, then loop

D) `range` over channels

`range ch` receives values until the channel is closed:

```go
for v := range ch {
	fmt.Println(v)
}
```

If the channel is never closed, the loop never ends.

------------------------------------------------------------

4) Go 1.22+ additions you used: `range` over integers

Starting Go 1.22, you can range over an integer `n` to iterate from `0` to `n-1`:

```go
for i := range 3 {
	fmt.Println(i) // 0, 1, 2
}
```

This is nice for small loops where you only need the index.

------------------------------------------------------------

5) Common pitfalls and best practices

Pitfall: taking the address of the range value

In `for _, v := range slice`, `v` is a loop variable. If you take `&v`, you’re taking the address of the loop variable, not the element.

Bad:

```go
vals := []int{1, 2, 3}
ptrs := []*int{}
for _, v := range vals {
	ptrs = append(ptrs, &v) // points to the same variable each iteration
}
```

Good (take the element address):

```go
for i := range vals {
	ptrs = append(ptrs, &vals[i])
}
```

Pitfall: relying on map order

If you need deterministic output (tests, logs, UI), sort keys.

Pitfall: modifying a slice while ranging over it

Appending to a slice while ranging can cause confusing behavior because the underlying array may reallocate.
Prefer:

- build a new slice
- or use an index loop if you really need in-place edits

Best practice: keep loops simple

- prefer early `continue` to reduce nesting
- prefer `range` when you want “iterate all items”
- prefer classic `for` when you need precise index control

------------------------------------------------------------

6) Patterns you’ll use often

A) “Repeat until success” with `for { ... }`

```go
for {
	err := doThing()
	if err == nil {
		break
	}
	time.Sleep(time.Second)
}
```

B) Loop with a timeout using `select`

```go
timeout := time.After(2 * time.Second)
for {
	select {
	case msg := <-ch:
		fmt.Println(msg)
	case <-timeout:
		return
	}
}
```

------------------------------------------------------------

Next experiments

- Extend `P6-Loops/main.go` with a `range` over a slice and a string, and print both index + value.
- Try a nested loop and use a label to break the outer loop.

