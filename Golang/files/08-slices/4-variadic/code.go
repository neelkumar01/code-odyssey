package main

import "fmt"

func myFunc(nums ...int) {
	fmt.Println(nums)
}

func sumAll(values ...int) int {
	sum := 0
	for i := 0; i < len(values); i++ {
		sum += values[i]
	}
	return sum
}

func printAll(things ...string) {
	for i := 0; i < len(things); i++ {
		fmt.Println(things[i])
	}
}

func main() {
	myFunc(1, 2, 3, 4, 5) // [1 2 3 4 5]

	/*
		myFunc() - variadic func
		‣ when we want to pass multiple args
		‣ all args passed got stored in slice
	*/
	netValue := sumAll(10, 20, 30)
	fmt.Println(netValue)

	/*
		spread operator
		‣ use to pass slice in variadic func
	*/
	basket := []string{"fruits", "veggies", "drink"}
	printAll(basket...)
}
