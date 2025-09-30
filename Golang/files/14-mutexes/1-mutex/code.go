package main

import (
	"fmt"
	"sync"
	"time"
)

var count = 0
func increment() {
	for i:=1; i<=1000; i++ {
		count += 1
	} 
}

var x = 0
var mux sync.Mutex	// creation of mutex
func update() {
	for i:=1; i<=1000; i++ {
		mux.Lock()
		x += 1
		mux.Unlock()
	}
}

func main() {

	// code with error

	go increment()
	go increment()
	time.Sleep(time.Second)
	fmt.Println("count:", count)

	// expected : 2000
	// output : random int (1609, 1500)
	// reason : both goroutines doing same thing at same time


	// fixed code with mutex
	// mutex : solve concurrent read/write problem
	go update()
	go update()
	time.Sleep(time.Second)
	fmt.Println("x:",x)
}