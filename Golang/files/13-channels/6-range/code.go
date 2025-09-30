package main

import "fmt"

func main() {
	ch := make(chan int, 3)	// buffered channel

	ch <- 1
	ch <- 2
	ch <- 3
	close(ch)

	for items := range ch {
		fmt.Println(items)
	}
}
