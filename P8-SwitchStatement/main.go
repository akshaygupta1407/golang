package main

import (
	"fmt"
	"time"
)

func main() {
	//simple switch
	i := 5

	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	case 4:
		fmt.Println("four")
	case 5:
		fmt.Println("five")
	default:
		fmt.Println("other")
	}

	//multiple condition switch
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It is weekend")
	default:
		fmt.Println("It is weekday")
	}

	//types switch
	whoAmI := func (i interface{})  { // if interface is empty that means it is any type
		switch t := i.(type) {
		case int:
			fmt.Println("It's an integer")
		case string:
			fmt.Println("It's an string")
		case bool:
			fmt.Println("It's a boolean")
		default:
			fmt.Println("It's a other type", t)
		}
	}

	whoAmI(24)
	whoAmI("golang")
	whoAmI(true)
	whoAmI(make(map[string]int))
}