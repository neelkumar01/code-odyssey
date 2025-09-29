package main

import "fmt"

func main() {
	// maps are like python dictionaries

	scores := make(map[string]int)

	scores["physics"] = 80
	scores["maths"] = 90
	scores["chemistry"] = 60

	fmt.Println(scores)

	ages := map[string]int{
		"neel": 18,
		"ram": 24,
	}
	fmt.Println(len(ages))
}