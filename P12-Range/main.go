package main

import "fmt"

//iterating over data structures
func main() {
	nums := []int{6, 7, 8}

	for idx, value := range nums {
		fmt.Println("Index:", idx, "value: ", value)
	}

	mpp := map[string]int{"key1": 1, "key2": 2}
	for key, value := range(mpp) {
		fmt.Println("key", key, "value", value)
	}


    //unicode code point rune
	str := "golang"
	for idx, ch := range str {
		fmt.Println("idx:", idx, "ch:", string(ch))
	}
}