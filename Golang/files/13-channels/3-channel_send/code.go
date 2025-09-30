package main

import (
	"fmt"
)

func worker(signal chan struct{}) {
	fmt.Println("working...")
	fmt.Println("done my work!")

	signal <- struct{}{}
}

func main() {
	signal := make(chan struct{})
	go worker(signal)

	<- signal  // blocks main to end untill something sent on channel
}