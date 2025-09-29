package main

import "fmt"

func square(n int) int {
	return  n * n
}

func transform(f func(int) int) func(int) int {
	return func(x int) int {
		return 2 * f(x)
	} 
}

func main() {
	/*
	Currying (a practice)
	‣ takes a func as arg in func
	‣ returns new func
	*/

	twoTimesSquare := transform(square)
	fmt.Println(twoTimesSquare(5))
}

