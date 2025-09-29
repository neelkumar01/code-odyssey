package main

import "fmt"

func createSlice(items ...int) []int {
	return items
}

func main() {
	// append(slice, ...things)
	// append is variadic function

	mySlice := createSlice(1,2,3,4)
	fmt.Println(mySlice)

	mySlice = append(mySlice, 5)
	mySlice = append(mySlice, 6,7)
	mySlice = append(mySlice, []int{8,9,10}...)

	fmt.Println(mySlice)
}
