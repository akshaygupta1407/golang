Maps in Go

A map in Go is a built-in hash table / dictionary / key-value store. It lets you look up values quickly by key, which makes it one of the most important data structures in everyday Go code.

This folder (`P11-Maps`) shows the core map patterns in `main.go`, and this README expands on how maps behave in Go.

From this example (`P11-Maps/main.go`)

Your program demonstrates:

- creating a map with `make`
- setting and reading values
- looking up missing keys
- checking map length
- deleting entries
- clearing a map
- literal map creation
- safe lookup with the `ok` idiom
- comparing maps with `maps.Equal`

Run it:

```bash
cd P11-Maps
go run .
```

------------------------------------------------------------

1) What a map is

A map stores data as key/value pairs.

Example:

```go
m := make(map[string]string)
```

This means:

- keys are `string`
- values are `string`

In general, a map type looks like:

```go
map[KeyType]ValueType
```

Maps are ideal when you want:

- fast lookup by key
- simple association between names and values
- a flexible dictionary-like structure

------------------------------------------------------------

2) Creating maps

Your example uses `make`:

```go
m := make(map[string]string)
```

You can also create a map with a literal:

```go
m := map[string]int{"price": 100}
```

Use `make` when you want to start with an empty map.
Use a literal when you already know the initial contents.

------------------------------------------------------------

3) Setting and getting values

To add or update an entry:

```go
m["name"] = "golang"
m["area"] = "backend"
```

To read an entry:

```go
fmt.Println(m["name"], m["area"])
```

Maps are indexed with square brackets, just like arrays and slices, but the index is a key instead of a numeric position.

------------------------------------------------------------

4) Zero value on lookup

Your code shows this:

```go
fmt.Println(m["djfn"])
```

If a key is missing, Go returns the zero value of the map’s value type.

For `map[string]string`, the zero value is `""`.
For `map[string]int`, the zero value is `0`.

Important:

- a missing key and a key whose value is the zero value can look the same if you only read the value
- use the `ok` idiom when you need to know whether the key actually exists

------------------------------------------------------------

5) The `ok` idiom

This is the safe way to check for presence:

```go
k, ok := m3["price"]
if ok {
	fmt.Println("All Ok", k)
}

k, ok = m3["age"]
if !ok {
	fmt.Println("All not Ok", k)
}
```

How it works:

- `k` receives the value
- `ok` is `true` if the key exists
- `ok` is `false` if the key is missing

This is one of the most important map patterns in Go.

------------------------------------------------------------

6) Length of a map

You can use `len` with maps:

```go
fmt.Println(len(m1))
```

For maps, `len` returns the number of key/value pairs currently stored.

------------------------------------------------------------

7) Deleting entries

To remove a key:

```go
delete(m1, "age")
```

Notes:

- deleting a missing key is safe
- after deletion, the key no longer exists in the map

This is useful for:

- removing stale data
- cleaning up state
- implementing caches or registries

------------------------------------------------------------

8) Clearing a map

Your example uses:

```go
clear(m1)
```

`clear` removes all entries from the map.

This is convenient when:

- you want to reuse the map variable
- you need to reset all entries quickly

After `clear`, the map is empty but still usable.

------------------------------------------------------------

9) Comparing maps

Go maps cannot be compared with `==` except against `nil`.

To compare contents, use `maps.Equal`:

```go
fmt.Println(maps.Equal(m2, m3))
```

This checks whether:

- the maps have the same keys
- the values for those keys are equal

This is the idiomatic comparison approach in modern Go.

------------------------------------------------------------

10) Map ordering

Map iteration order is not guaranteed in Go.

That means:

- you should not rely on the order of keys when printing or ranging over a map
- if you need stable output, collect and sort the keys first

This is a common source of confusion when people first work with maps.

------------------------------------------------------------

11) Common map patterns

A) Configuration or settings

```go
settings := map[string]string{
	"env": "dev",
	"region": "asia",
}
```

B) Counting

```go
counts := map[string]int{}
counts["go"]++
```

C) Existence checks

```go
if _, ok := users[id]; ok {
	fmt.Println("user exists")
}
```

D) Set-like behavior

```go
visited := map[string]bool{}
visited["a"] = true
```

------------------------------------------------------------

12) Best practices

- use maps when lookup by key matters
- use the `ok` idiom when missing keys are meaningful
- prefer literals when the initial data is known
- use `delete` and `clear` for cleanup
- use `maps.Equal` instead of trying to compare maps directly
- remember that map iteration order is random

------------------------------------------------------------

Next experiments

- Add a `range` loop over `m` and print keys and values.
- Try a `map[string]bool` to represent a set.
- Sort map keys before printing them for stable output.

