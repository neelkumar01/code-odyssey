package nil_Pointers
import "fmt"

var mySlice []int
var m map[string]int

func Nil() {
	length := len(mySlice)		// length = 0
	val, ok := m["name"]

	fmt.Println(length)
	fmt.Printf("value of [\"name\"] = %v  ok = %v", val, ok)

	
}