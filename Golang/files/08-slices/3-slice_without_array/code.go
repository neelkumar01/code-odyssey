package main

import "fmt"

func main() {
	// using make
	// make([]type, len, cap)

	mySlice := make([]string, 3, 5)
	// length : elements in slice
	// capacity : elements in underlying array

	mySlice[0] = "mango"
	mySlice[1] = "strawberry"
	mySlice[2] = "blueberry"
	fmt.Println(mySlice)

	// another way
	veggies := []string{"tomato", "capsicum"}
	fmt.Println(veggies)
}