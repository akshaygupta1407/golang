package main

import (
	"fmt"
	"maps"
)

// maps -> hash, object, dict
func main() {

	//creating map
	m := make(map[string]string)

	//setting an element
	m["name"] = "golang"
	m["area"] = "backend"

	//getting an element
	fmt.Println(m["name"], m["area"]) //golang backend
	fmt.Println(m)                    //map[area:backend name:golang]
	fmt.Println(m["djfn"])            //returned zeroes value

	m1 := make(map[string]int)
	m1["age"] = 30
	fmt.Println(m1["age"]) //30
	fmt.Println(len(m1))   //1

	m1["price"] = 200
	fmt.Println(m1) //map[age:30 price:200]
	delete(m1, "age")
	fmt.Println(m1) //map[price:200]

	clear(m1)
	fmt.Println(m1) // map[]

	m2 := map[string]int{"price": 100}
	fmt.Println(m2) //map[price:100]

	m3 := map[string]int{"price": 100}
	k, ok := m3["price"]
	if ok {
		fmt.Println("All Ok", k) //All Ok 100
	}
	k, ok = m3["age"]
	if !ok {
		fmt.Println("All not Ok", k) //All not Ok 0
	}

	fmt.Println(maps.Equal(m2, m3)) //true

}
