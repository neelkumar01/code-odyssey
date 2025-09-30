package main

import "fmt"

func send(ch chan string) {
	ch <- "msg from neel"
}

func receive(ch chan string) {
	msg := <-ch
	fmt.Println(msg)
}

func main() {
	// channels - allow goroutines to communicate
	ch := make(chan string)

	go send(ch)
	go receive(ch)

	// enter any num to exit
	var x int;
	fmt.Scanln(&x)
}