package main

import "fmt"

type car struct {
	color string
}
func (c *car) setColor(color string) {
	c.color = color
}

func main() {
	c := car{
		color: "blue",
	}
	c.setColor("black")
	fmt.Println(c.color)
}