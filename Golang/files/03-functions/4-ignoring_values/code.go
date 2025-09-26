package main

import "fmt"

func main() {
	// ignored val of y using underscore ( _ )
	x, _ := getOriginCoordinates()
	fmt.Println(x)
}

func getOriginCoordinates() (x , y int) {
	return 0, 0
}

