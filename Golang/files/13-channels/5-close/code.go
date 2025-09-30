package main

import "fmt"

func main() {

	ch := make(chan int)

	close(ch)	// closing channel

	// to check channel status
	_, ok := <-ch	// consit of value at _
	if ok {
		fmt.Println("channel is open")
	} else {
		fmt.Println("channel is closed!")
	}
}