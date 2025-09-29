package main

import "fmt"

func main() {
	warehouse := make(map[string]int)  // product category : items available

	// insert
	warehouse["TV"] = 40
	fmt.Println(warehouse)
	
	// get elem
	elem := warehouse["TV"]
	fmt.Println(elem)

	// delete elem
	delete(warehouse, "TV")
	fmt.Println(warehouse)

	// to check if elem exist
	value, ok := warehouse["TV"]
	fmt.Println(value, ok)
}