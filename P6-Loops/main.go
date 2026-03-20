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

	//infinite loop
	/*
		for {
			fmt.Println(1)
		}
	*/

	fmt.Println("classic loop")
	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}

	//classic loop with break
	fmt.Println("classic loop with break")
	for i := 0; i < 3; i++ {
		fmt.Println(i)
		if i == 1 {
			fmt.Println("Breaking after index", i)
			break
		}
	}

	//in go 1.22 Range
	fmt.Println("Range")
	for i:= range 3 {
		fmt.Println(i)
	}

}
