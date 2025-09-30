package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		// delay current goroutine for 2s
		ch1 <- "msg1"
	}()

	go func() {
		time.Sleep(1 * time.Second)
		// delay current goroutine for 1s
		ch2 <- "msg2"
	}()

	// receiving channel data
	select {
	case msg1 := <-ch1:
		fmt.Println(msg1)
	case msg2 := <-ch2:
		fmt.Println(msg2)
	default:
		fmt.Println("no channel ready")
	}
}
