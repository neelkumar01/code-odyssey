package main

import "fmt"

/*
Methods
- these are functions
- they have a receiver
*/

type rect struct {
	width int
	height int
}

func (r rect) area() int {
	return r.width * r.height
}
/*
area() - name of method
r react - receiver (to which our method associated)
*/

func main() {
	box := rect{
		width: 10,
		height: 20,
	}

	fmt.Println(box.area())
}
