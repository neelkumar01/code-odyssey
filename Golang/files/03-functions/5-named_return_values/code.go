package main

import "fmt"

func main() {
	x, y := getCoord()
	fmt.Println(x, y)	// 0 0

	a, b := getPoint()
	fmt.Println(a, b)
}

func getCoord() (x,y int) {
	return	// this directly returs x and y
}

func getPoint() (int, int) {
	var x int
	var y int
	return  x, y
}

// both getCoord() annd getPoint() func are same !