package main

import (
	"fmt"
	"strconv"
)

func main() {
	// error handling
	
	value, err := strconv.Atoi("12a")  // str to int
	if err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Println(value)
	}

	value2, err2 := strconv.Atoi("12")  
	if err2 != nil {
		fmt.Println("error:", err2)
	} else {
		fmt.Println(value2)
	}
}