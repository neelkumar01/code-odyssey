package main

import "fmt"

func add(x,y int) int {
	return x + y
}

func calculate(a,b,c int, solve func(int, int) int) int {
		return solve(solve(a,b), c)
	}

func main() {
	/*
	first class func - when func treated as variables
	higher order func - accept or returns a function
	*/
	answer := calculate(1,2,3, add)
	fmt.Println(answer)
}