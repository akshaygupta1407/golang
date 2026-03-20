package main

import "fmt"

// for -> only construct in go for looping
// Note: no while keyword in Go
func main() {

	//while loop
	fmt.Println("while loop")
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i++
	}

	fmt.Println("######################")

	//infinite loop
	/*
		for {
			fmt.Println(1)
		}

	fmt.Println("######################")
	*/

	fmt.Println("classic loop")
	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}

	fmt.Println("######################")

	//classic loop with break
	fmt.Println("classic loop with break")
	for i := 0; i < 3; i++ {
		fmt.Println(i)
		if i == 1 {
			fmt.Println("Breaking after index", i)
			break
		}
	}

	fmt.Println("######################")

	//in go 1.22 Range
	fmt.Println("Range")
	for i := range 3 {
		fmt.Println(i)
	}

	fmt.Println("######################")

	fmt.Println("iterate over a collection / string / channel / integer")
	nums := []int{1 ,2, 3}
	for idx, value := range(nums) {
		fmt.Println("Index & Value: ", idx, value)
	}

	fmt.Println("######################")

	s := "hé"
	for idx, r := range s {
		fmt.Println(idx, r) // idx is byte index; r is a rune
	}

	fmt.Println("######################")

	fmt.Println("Range in Maps")

	mpp := map[string]int{"a":1, "b":2}
	for k,v := range(mpp) {
		fmt.Println("Key, value: ", k, v)
	}

	fmt.Println("######################")


	fmt.Println("Labels")
	outer:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Println("i & j", i, j)
			if i == 1 && j == 1 {
				break outer
			}
		}
	}

	fmt.Println("######################")



}
