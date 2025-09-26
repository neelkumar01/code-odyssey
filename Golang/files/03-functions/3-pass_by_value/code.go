package main

import "fmt"

func main() {
	carFuel := 12
	addFuel(carFuel)
	fmt.Println(carFuel)	// 12 not 13
}

func addFuel(fuel int) int {
	return fuel+1
}
