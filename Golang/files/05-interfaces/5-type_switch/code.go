package main

import (
	"fmt"
)

/*
func myFunc(data interface{}) {}
‣ interface{} is empty interface
‣ don't have any methods in it
‣ macthes to every type in go
*/

// previous example
func myFunc(arg interface{}) {
	// check if type of arg is string and likewise give bool to ok
	value, ok := arg.(string) 
	if ok {
		fmt.Println(value)
	} else {
		fmt.Println("not a string!")
	}
}

// switch case!
func myFunc2(arg interface{}) {
	switch value := arg.(type) {
	case int:
		fmt.Printf("%v, %T \n", value, value)
	case string:
		fmt.Printf("%v, %T \n", value, value)
	default:
		fmt.Printf("%v, %T \n", value, value)
	}
}

func main() {
	myFunc("neel")
	myFunc(12)

	myFunc2(10)
}