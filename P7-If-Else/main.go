package main

import "fmt"

func main() {
	age := 16
	if age >= 18 {
		fmt.Println("person is an adult")
	} else if age >= 12 {
		fmt.Println("person is a teenager")
	} else {
		fmt.Println("person is a kid")
	}

	var role = "admin"
	var hasPermission = true

	if role == "admin" || hasPermission {
		fmt.Println("yes user is have access")
	}

	if role == "admin" && hasPermission {
		fmt.Println("yes user is have access")
	}

	//we can declare a variable inside if construct
	if age:=15; age >= 18 {
		fmt.Println("person is an adult. age: ", age)
	} else if age >= 12 {
		fmt.Println("person is a teenager. age: ", age)
	}

	// go does nto have ternary operator. You will have to use nornal if else
}