package main

import "fmt"

type car struct {
	brand string
	color string
	seats int

	frontwheel wheel
	backwheel wheel
}
type wheel struct {
	radius int
	material string
}

func main() {
	myCar := car{}
	myCar.brand = "rolls royce"
	fmt.Println(myCar.brand)

	myCar.frontwheel.radius = 10
	fmt.Println(myCar)
}