package maps
import "fmt"

var myMap = map[string]int {
	"one": 1,
	"two": 2,
	"three":3,
}

var nilMap map[string]int

func Literal() {
	var1 := nilMap == nil
	fmt.Println(var1)		// true

	fmt.Println(len(myMap))
}