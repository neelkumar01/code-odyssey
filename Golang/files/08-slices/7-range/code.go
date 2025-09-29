package main

import "fmt"

func main() {

	mySlice := []string{"mango", "pineapple", "blueberry"}

	for index, element := range mySlice {
		fmt.Println(index, element)
	}
}
