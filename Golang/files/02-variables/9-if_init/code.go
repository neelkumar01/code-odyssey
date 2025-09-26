package main

import "fmt"

func main() {
	// declaring initial statements in if block only
	if age := 18; age < 18 {
		fmt.Println("you can't drive")
	} else {
		fmt.Println("you can drive")
	}
}