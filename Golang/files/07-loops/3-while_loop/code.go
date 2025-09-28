package main

import "fmt"

func main() {
	/*
	‣ There is no while loop
	‣ for loop acts as while loop
	‣ by using concept of omission
	*/
	plantHeight := 1

	for plantHeight <= 5 {
		fmt.Printf("Plant is growing...\nPlant height: %v\n", plantHeight)

		plantHeight++
	}
	fmt.Printf("Plant has grown %vm tall!", plantHeight-1)
}