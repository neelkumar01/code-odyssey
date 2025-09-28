package main

import "fmt"

func sendingCost(msgCount int) float64 {
	totalCost := 0.0
	costPerMsg := 1.0
	for i:=0; i<msgCount; i++ {
		totalCost += costPerMsg
	}
	return totalCost
}
func sendAll(msgCount int) {
	fmt.Printf("Sending %v messages...\n", msgCount)
	totalCost := sendingCost(msgCount)
	fmt.Printf("All msgs sent at cost of Rs.%v", totalCost)
}
func main() {
	// sending msgs in bulk
	sendAll(12)
}