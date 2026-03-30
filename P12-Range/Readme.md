Range in Go

`range` is Go’s general-purpose iteration tool. It lets you loop over slices, arrays, maps, strings, channels, and even integers in modern Go. In practice, it is one of the most frequently used language features because it keeps iteration short and readable.

This folder (`P12-Range`) shows the core `range` patterns in `main.go`, and this README adds a deeper theory section on how `range` behaves in different situations.

From this example (`P12-Range/main.go`)

Your program demonstrates:

- iterating over a slice
- iterating over a map
- iterating over a string
- receiving index/value pairs from `range`
- rune-aware string iteration

Run it:

```bash
cd P12-Range
go run .
```

------------------------------------------------------------

1) What `range` does

`range` is not a separate data structure. It is syntax for iterating over a collection.

General pattern:

```go
for key, value := range collection {
	fmt.Println(key, value)
}
```

What the two variables mean depends on the type being ranged over.

------------------------------------------------------------

2) `range` over slices and arrays

Your example starts with a slice:

```go
nums := []int{6, 7, 8}

for idx, value := range nums {
	fmt.Println("Index:", idx, "value:", value)
}
```

For slices and arrays:

- the first value is the index
- the second value is the element at that index

Important details:

- the index starts at `0`
- the loop runs in order from left to right
- the value is a copy of the element, not the element itself

If you only need the index:

```go
for idx := range nums {
	fmt.Println(idx)
}
```

If you only need the values:

```go
for _, value := range nums {
	fmt.Println(value)
}
```

This is one of the most common idioms in Go.

------------------------------------------------------------

3) `range` over maps

Your example also ranges over a map:

```go
mpp := map[string]int{"key1": 1, "key2": 2}

for key, value := range mpp {
	fmt.Println("key", key, "value", value)
}
```

For maps:

- the first value is the key
- the second value is the value stored at that key

Important map behavior:

- iteration order is not guaranteed
- the order can change between runs
- you should not rely on map order for logic or tests

If you only need keys:

```go
for key := range mpp {
	fmt.Println(key)
}
```

If you need stable output, collect the keys, sort them, then iterate.

------------------------------------------------------------

4) `range` over strings

Your code ranges over a string:

```go
str := "golang"

for idx, ch := range str {
	fmt.Println("idx:", idx, "ch:", string(ch))
}
```

This is a very important Go detail.

For strings, `range` decodes UTF-8 and gives you:

- the byte index
- the rune value

That means `ch` is a `rune`, not a `byte`.

Why that matters:

- ASCII characters are one byte and one rune
- many Unicode characters take multiple bytes
- `range` gives you whole characters, not raw bytes

If you want raw bytes instead:

```go
for i := 0; i < len(str); i++ {
	fmt.Println(str[i])
}
```

If you want Unicode-aware character iteration, `range` is the right tool.

------------------------------------------------------------

5) Byte index vs character index

The index returned by `range` on a string is a byte index, not a character count.

For example, with a multi-byte character like `é`:

```go
for idx, ch := range "hé" {
	fmt.Println(idx, ch)
}
```

The index jumps by the byte length of each rune.

That is why string `range` is great for Unicode, but not for “character count” unless you count the iterations yourself.

------------------------------------------------------------

6) Range variable behavior

Inside a `range` loop, the loop variables are reused on each iteration.

This matters when you take addresses or capture variables in closures.

Example pitfall:

```go
for _, value := range nums {
	fmt.Println(&value)
}
```

`value` is the loop variable, not the original element.

If you need the actual element address, use the index:

```go
for i := range nums {
	fmt.Println(&nums[i])
}
```

This is a common Go gotcha.

------------------------------------------------------------

7) `range` over channels

Although your sample does not include a channel, `range` also works with channels:

```go
for v := range ch {
	fmt.Println(v)
}
```

This keeps receiving until the channel is closed.

This is often used in worker pipelines and producer/consumer code.

------------------------------------------------------------

8) `range` over integers in Go 1.22+

Modern Go also allows:

```go
for i := range 3 {
	fmt.Println(i)
}
```

That iterates `0, 1, 2`.

This is useful when:

- you only need an index
- you want a short loop for a known count
- you do not need to declare and manage the counter manually

------------------------------------------------------------

9) Choosing between `range` and classic `for`

Use `range` when:

- you want to iterate over every item
- you are working with slices, maps, strings, or channels
- readability matters more than low-level control

Use classic `for` when:

- you need custom stepping
- you need to iterate backward
- you need exact control over the loop counter

------------------------------------------------------------

10) Common pitfalls

- map order is random, so do not depend on it
- string indices from `range` are byte offsets, not character counts
- loop variables are reused, so be careful with addresses and closures
- if you only need one of the two `range` results, use `_` for the one you do not need

------------------------------------------------------------

11) Best practices

- use `range` for most everyday iteration
- use `_` to ignore values you do not need
- use indices when you need direct access to slice elements
- use `range` on strings for Unicode-aware iteration
- sort map keys when output order matters

------------------------------------------------------------

Next experiments

- Add a `range` loop over a map with sorted keys for stable output.
- Try a string with multi-byte Unicode characters and print the byte index and rune.
- Compare `for i := range nums` with `for idx, value := range nums`.

