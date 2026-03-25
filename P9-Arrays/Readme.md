Arrays in Go

An array in Go is a numbered sequence with a fixed length. That fixed size is part of the type, which makes arrays different from slices.

This folder (`P9-Arrays`) shows the core array patterns in `main.go`, and this README adds a deeper explanation of how arrays behave in Go.

From this example (`P9-Arrays/main.go`)

Your program demonstrates:

- declaring an array with a fixed length
- zero values for array elements
- assigning values by index
- array literals
- a 2D array

Run it:

```bash
cd P9-Arrays
go run .
```

------------------------------------------------------------

1) What makes an array an array

In Go, an array has:

- a fixed length
- a single element type
- contiguous storage

Example:

```go
var nums [4]int
```

This means:

- the array always has exactly 4 elements
- each element is an `int`
- the length is part of the type: `[4]int`

Because the size is fixed, `[4]int` and `[5]int` are different types.

------------------------------------------------------------

2) Zero values

When you declare an array without initializing it, Go fills it with zero values.

```go
var nums [4]int
fmt.Println(nums[0])
```

For `int`, the zero value is `0`.

Other common zero values:

- `string` -> `""`
- `bool` -> `false`
- pointers, maps, slices, channels, functions -> `nil`

This is why a freshly declared array prints like:

```go
[0 0 0 0]
```

------------------------------------------------------------

3) Indexing and assignment

Arrays are accessed by index, starting at `0`.

```go
nums[0] = 1
nums[1] = 2
nums[2] = 3
nums[3] = 4
```

Key points:

- the first element is `nums[0]`
- the last element in a 4-element array is `nums[3]`
- accessing an invalid index causes a runtime panic

Example:

```go
fmt.Println(nums)
```

This prints the full array value after assignment.

------------------------------------------------------------

4) Declaring arrays in one line

You can initialize an array with a literal:

```go
nums1 := [4]int{0, 1, 2, 3}
```

This is useful when:

- the values are known up front
- you want the array contents to be obvious at a glance
- the array is small and fixed

You can also let Go infer the length:

```go
nums := [...]int{10, 20, 30}
```

Here, Go counts the elements for you.

------------------------------------------------------------

5) Partially initialized arrays

You can set only some elements and leave the rest at their zero values.

```go
var names [3]string
names[1] = "golang"
fmt.Println(names)
```

The result is:

```go
[  golang ]
```

The missing elements stay as empty strings.

This is useful when:

- only a few positions matter
- you want to update selected slots by index

------------------------------------------------------------

6) 2D arrays

Your example also includes a 2D array:

```go
nums2 := [4][2]int{{0, 1}, {2, 3}, {4, 5}, {6, 7}}
```

That means:

- 4 rows
- each row contains 2 integers

You can think of it as a fixed-size table:

```go
[ [0 1]
  [2 3]
  [4 5]
  [6 7] ]
```

2D arrays are useful for:

- matrix-style data
- fixed grids
- tabular data with a known shape

------------------------------------------------------------

7) Arrays vs slices

Arrays are less common in everyday Go than slices, but they still matter.

Use arrays when:

- the size is known and truly fixed
- you want the length to be part of the type
- you are working with low-level, fixed-size data

Use slices when:

- the size may change
- you want a more flexible collection
- you are building lists, buffers, or dynamic data

Important difference:

- arrays are values
- slices are descriptors over an underlying array

That means arrays are copied when assigned or passed around, which can be expensive for larger arrays.

------------------------------------------------------------

8) Common pitfalls

- array lengths are part of the type, so `[3]int` and `[4]int` are not interchangeable
- indexing past the end panics
- arrays are copied by value, so changes to a copied array do not affect the original
- for large collections, slices are usually the better default

------------------------------------------------------------

9) Best practices

- use arrays for fixed-size data with a known length
- use literals when the full contents are known upfront
- prefer slices for most application-level collection work
- keep 2D arrays for truly fixed shapes, such as small matrices or grids

------------------------------------------------------------

Next experiments

- Add a `range` loop over `nums` and print index plus value.
- Compare assigning an array to another variable versus assigning a slice.
- Try a 2D array with row and column iteration.

