package main

import "fmt"

const age = 30
// name1:="fjf"   we cant write shorthand variable declaration at global level

func main() {
	const name string = "Hello"
	//name = "m"    cannot assign to name (neither addressable nor a map index expression) 
	fmt.Println(name)


	//constant grouping
	const (
		host = "localhost" //can declare type also like host string
		port = 5000
	)

	fmt.Println(host, port)
}