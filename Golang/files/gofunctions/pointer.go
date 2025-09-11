package gofunctions
import "fmt"

func Pointer() {
	x := 12
	p := &x
	// p is a pointer : stores memory address of other variables
	fmt.Println("address of x :", p)

	// derefrencing with pointer
	fmt.Println("original x :", x)
	*p = 24
	fmt.Println("new val of x :", x)
}