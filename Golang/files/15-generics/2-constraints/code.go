package main

import "fmt"

type num interface {
	int | float64 // This means num can be int or float64
}

func add[T num](a, b T) T {
	return a + b
}

func main() {

	// constraint - when we want a specific type
	fmt.Println(add(1,2.4))
}
