package main

import (
	"fmt"
)

func main() {
	var username string = "wagslane"
	var password string = "2094738282"
	fmt.Println("Authorization: Basic", username+":"+password)

}