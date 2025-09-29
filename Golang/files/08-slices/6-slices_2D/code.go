package main

import "fmt"

func main() {
	studentInfo := [][]int{
		{1, 23},  // person 1
		{2, 12},  // person 2
		{3, 45},  // person 3
	}
	// i - index of each element of outer slice
	// student - each element in outer slice
	for i, person := range studentInfo {

		fmt.Printf("Details of student %d \n", i+1)

		fmt.Printf("id : %v \n", person[0])
		fmt.Printf("seat number : %v \n\n", person[1])
	} 
}