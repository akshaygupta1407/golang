Slices in Go

A slice is Go’s flexible, dynamic view over an underlying array. In everyday Go, slices are used far more often than arrays because they grow, shrink, and make collection code much easier to work with.

This folder (`P10-Slices`) shows the core slice patterns in `main.go`, and this README adds a deeper explanation of how slices behave in Go.

From this example (`P10-Slices/main.go`)

Your program demonstrates:

- a nil slice
- creating slices with `make`
- length vs capacity
- `append`
- an empty but non-nil slice
- `copy`
- slicing syntax
- slice comparison with `slices.Equal`
- 2D slices

Run it:

```bash
cd P10-Slices
go run .
```

------------------------------------------------------------

1) What a slice is

A slice is not the array itself. It is a small descriptor that points to an underlying array and carries:

- a pointer to the backing array
- a length
- a capacity

That is why slices can grow and shrink without you manually managing storage.

Example:

```go
var nums []int
```

This declares a slice of `int` with no backing array yet.

------------------------------------------------------------

2) Nil slices

An uninitialized slice is `nil`:

```go
var nums []int
fmt.Println(nums == nil)
fmt.Println(len(nums))
```

For a nil slice:

- `nums == nil` is `true`
- `len(nums)` is `0`
- `cap(nums)` is `0`

Nil slices are common and perfectly valid. They behave like empty slices in many situations, but they are not the same thing.

------------------------------------------------------------

3) Creating slices with `make`

`make` allocates a slice with a given length and optionally capacity:

```go
nums1 := make([]int, 2)
nums3 := make([]int, 3, 5)
```

The parameters mean:

- first argument: slice type
- second argument: length
- third argument: capacity, optional

Important difference:

- length is how many elements are currently in the slice
- capacity is how many elements can fit before reallocation is needed

Example:

```go
fmt.Println(len(nums3))
fmt.Println(cap(nums3))
```

------------------------------------------------------------

4) Length vs capacity

In your example:

```go
var nums3 = make([]int, 3, 5)
```

This creates:

- length 3
- capacity 5

That means the slice already has 3 accessible elements, and it can grow to 5 before Go needs to allocate a new backing array.

This is why you can append a few values before the capacity changes.

------------------------------------------------------------

5) `append`

`append` adds elements to the end of a slice:

```go
nums3 = append(nums3, 1)
```

If there is spare capacity, Go reuses the same backing array.
If capacity is exhausted, Go allocates a new backing array with more room.

That is why capacity can change after repeated appends.

Important rule:

- always assign the result of `append` back to the slice variable

Example:

```go
nums3 = append(nums3, 2)
nums3 = append(nums3, 3)
```

------------------------------------------------------------

6) Empty slice vs nil slice

Your example also shows:

```go
nums4 := []int{}
fmt.Println(nums4 == nil)
fmt.Println(nums4)
fmt.Println(len(nums4))
fmt.Println(cap(nums4))
```

This is an empty slice literal.

Key difference:

- `[]int{}` is empty but not nil
- `var nums []int` is nil

Both have length `0`, but they differ in equality checks and sometimes in API behavior.

------------------------------------------------------------

7) Copying slices with `copy`

Go provides `copy` to duplicate slice elements:

```go
nums5 := make([]int, 0, 5)
nums5 = append(nums5, 1)
nums6 := make([]int, len(nums5), 5)
copy(nums6, nums5)
```

Important notes:

- `copy(dst, src)` copies elements, not the slice header
- it copies up to the smaller of the two lengths
- destination length controls how many elements can be received

This is useful when:

- you want a separate backing array
- you need to preserve original values
- you want to safely modify a clone

------------------------------------------------------------

8) Slicing syntax

Slices can be carved out of existing slices:

```go
nums7 := []int{1, 2, 3}
fmt.Println(nums7[0:2])
fmt.Println(nums7[1:])
```

The general form is:

```go
slice[low:high]
```

Rules:

- `low` is inclusive
- `high` is exclusive
- omitting `low` means start at `0`
- omitting `high` means go to the end

Important:

- slicing does not copy the data
- the new slice shares the same underlying array

That means changing the new slice may affect the original slice if they share storage.

------------------------------------------------------------

9) Equality and `slices.Equal`

Slices cannot be compared with `==` except against `nil`.

To compare contents, use `slices.Equal`:

```go
nums8 := []int{1, 2}
nums9 := []int{1, 2}
fmt.Println(slices.Equal(nums8, nums9))
```

This returns `true` when:

- lengths match
- corresponding elements are equal

This is the idiomatic way to compare slice contents in modern Go.

------------------------------------------------------------

10) 2D slices

Your example includes a 2D slice:

```go
nums10 := [][]int{{1, 2}, {3, 4}}
```

This is a slice of slices.

Why use this instead of a 2D array?

- rows can vary in length
- it is more flexible
- it is common for table-like or jagged data

You can think of it as:

- an outer slice containing rows
- each row is another slice

------------------------------------------------------------

11) Arrays vs slices

Arrays are fixed-size values.
Slices are dynamic views over arrays.

Use slices when:

- the collection changes size
- you want flexible APIs
- you are handling lists, buffers, or table rows

Use arrays when:

- the size is fixed and important
- you need the length in the type
- you are working with small, rigid structures

In most application code, slices are the default choice.

------------------------------------------------------------

12) Best practices

- prefer slices for most collection work
- use `make` when you know the size or want to reserve capacity
- remember that `append` may reallocate
- use `copy` when you need an independent backing array
- use `slices.Equal` for content comparison
- use 2D slices when row lengths may vary

------------------------------------------------------------

Next experiments

- Print `len` and `cap` after each `append` to watch growth.
- Try modifying a subslice and see how it affects the original slice.
- Compare a nil slice and an empty slice in a small example.

