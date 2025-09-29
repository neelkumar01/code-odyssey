package main

import "fmt"

func counter() func() int {
	count := 0
	return func() int {
		count ++
		return count
	}
}

func main() {

	// closure = inner function + capture variables

	increment := counter()

	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())
}