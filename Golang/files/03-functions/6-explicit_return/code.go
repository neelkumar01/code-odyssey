package main

import "fmt"

func main() {
	x, y := getPoints()
	fmt.Println(x, y)
}

func getPoints() (x, y int) {
	return 5, 8
}