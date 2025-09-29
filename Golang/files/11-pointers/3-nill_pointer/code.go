package main

import "fmt"

func main() {

	// if pointer dont point anything - cause runtime panic error

	// handling nil pointers
	var pointer *int   // uninitialised pointer

	// fmt.Println(*pointer) - give panic runtime error

	if pointer == nil {
		fmt.Println("pointer is nill")
	} else {
		fmt.Println(*pointer)
	}

	// initialising pointer
	value := 24
	pointer = &value
	fmt.Println(*pointer)
}