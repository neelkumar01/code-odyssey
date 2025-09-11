package goTypes
import "fmt"
func StructuralTyping() {
	a := [...]int{1,2,3}		// [1, 2, 3]
	b := [3]int{}				//  [0, 0, 0]

	a = b		// OK

	c := [4]int{}
	// a = c 		--->   not OK

	fmt.Println(a,c)
}