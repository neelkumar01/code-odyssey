package main

import "fmt"

func main() {
	const secondsInMinutes = 60
	const minutesInHour = 60
	// computed const :-
	const secondsInhour = secondsInMinutes * minutesInHour
	fmt.Println(secondsInhour)
}

