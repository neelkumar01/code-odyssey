package maps
import "fmt"

func Index() {
	myMap1 := map[string]int{}		// non-nil but empty
	a := myMap1["key"]		// return 0
	fmt.Println(a)

	b, ok := myMap1["key"]			// ok tells whether key in map or not
	fmt.Printf("b = %v   ok = %v \n", b, ok)

	myMap1["key"]++

	c, ok := myMap1["key"]
	fmt.Printf("c = %v  ok = %v \n", c, ok)
}