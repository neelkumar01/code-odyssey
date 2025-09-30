package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	note   = "hello"
	rwLock sync.RWMutex // read-write lock
)

func read(id int) {

	rwLock.RLock()   // read lock

	fmt.Printf("Reader %v is reading: %s \n", id, note)

	time.Sleep(time.Second)   // allowing goroutine to run completely

	rwLock.RUnlock()
}

func write(data string) {

	rwLock.Lock()	// write lock

	note = data

	time.Sleep(time.Second)

	rwLock.Unlock()
}

func main() {

	go read(1)		// reader 1 
	go read(2)		// reader 2

	time.Sleep(time.Second)

	go write("new data")

	time.Sleep(time.Second)

	go read(3)

	time.Sleep(time.Second)

}

/*
	• RLock 
		‣ multiple goroutines can read
		‣ if anyone is writing, put a LOCK!
		‣ when any goroutine write, other can't even read
*/
