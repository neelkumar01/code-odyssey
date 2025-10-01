package main

import "fmt"

func print[T any](value T) {
	fmt.Println(value)
}

func main() {

	/*
	• Earlier interface{} used to accept any type
	• Go v1.18 introduced generics 
		‣ more type safe
		‣ work with any type
		‣ reduce repetitive code
	*/

	print("neel")	// string
	print(1)		// int
	print(2.4)		// float
}