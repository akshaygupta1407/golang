package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func getLanguages() (string, string, string) {
	return "Go", "Python", "JavaScript"
}

func main() {
	sum := add(3, 5)
	fmt.Println(sum)
	lang1, lang2, lang3 := getLanguages()
	fmt.Println(lang1, lang2, lang3)
}