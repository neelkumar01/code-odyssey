package main

import "fmt"

func main() {
	// printing odd numbers
	fmt.Println("printing odd numbers")
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}

	// using break statement
	fmt.Println("using break")
	for k := 0; k <= 10; k++ {
		if k == 6 {
			break
		}
		fmt.Println(k)
	}
}
