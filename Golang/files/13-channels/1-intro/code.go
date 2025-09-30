package main

import (
	"fmt"
	"time"
)

func sayHi() {
	fmt.Println("Hello")
}

func main() {
	// concurrency - doing multiple tasks at one time

	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	go sayHi()  // goroutine - run in bg

	time.Sleep(1 * time.Second) // pauses main function 
	// so that when all code got executed, goroutines can also execute after them
}

