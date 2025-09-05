package maps
import "fmt"

func Index() {
	myMap1 := map[string]int{}		// non-nil but empty
	a := myMap1["key"]		// return 0
	fmt.Println(a)

	b, ok := myMap1["key"]
	fmt.Printf("b = %v   ok = %v", b, ok)

	myMap1["key"]++
}