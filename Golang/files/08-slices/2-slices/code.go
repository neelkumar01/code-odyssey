package main

import "fmt"

func main() {
	/*
	• 99% time - slices are used over array
	• slice size : dynamic, flexible
	• array size : fixed
	*/

	myArr := [3]int{1,2,3}
	mySlice := myArr[0:2]	// [1 2]
	fmt.Println(mySlice)


}