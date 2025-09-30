package main

import "fmt"

func main() {
	ch := make(chan int, 2)

	// buffer channel - stores limited data
	
	ch <- 10
	ch <- 20

	fmt.Println(<-ch)
	fmt.Println(<-ch)
}