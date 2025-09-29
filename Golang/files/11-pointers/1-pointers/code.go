package main

import "fmt"

func main() {

	// variable - name of memory location where data is stored
	x := 24
	// location name : x
	// data stored : 24

	// pointer - stores moemory address of variable
	// so point directly to location!

	address := &x
	fmt.Println(address)

	// pointer helps to chnage data directly in memory

	// reading data using pointer
	data := "unknown data here..."
	dataPtr := &data
	fmt.Println(*dataPtr)	// dereferencing

	*dataPtr = "This data got changed!"
	fmt.Println(data)
}