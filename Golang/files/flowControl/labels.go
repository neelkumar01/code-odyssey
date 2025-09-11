package flowControl
import "fmt"

func Label() {
	myMap := map[string]int{
		"one":1,
		"two":2,
		"three":3,
		"four":4,
		"five":5,
	}
	slicedData := []string{
		"one",
		"two",
		"three",
	}

	//   _   :   this is a blank identifier
	outerLoop:			// outerLoop  --->  label
		for mapKey := range myMap {
			for _, sliceVal := range slicedData { 
				if mapKey == sliceVal {
					fmt.Println("found :", mapKey)
					continue outerLoop
				}
			}
		}
}

