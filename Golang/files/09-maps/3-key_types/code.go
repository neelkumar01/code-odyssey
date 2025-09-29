package main

func main() {
	/*
	‣ only comparable type can be made keys
	‣ keys can't be slice, function, map
		‣ these can't be compared with ==, !=
	*/

	// to map how many times a airline comes from a city
	// Map: city => airline => count

	// using struct
	type Tracker struct {
		city 	string
		airline string
	}

	arrivalData := make(map[Tracker]int)

	arrivalData[Tracker{"delhi", "indigo"}]++
	arrivalData[Tracker{"delhi", "air india"}]++
	arrivalData[Tracker{"banglore", "indigo"}]++
}