package slice
import "fmt"

/*
slice  --->  stores array's info and address
slice also known as descriptor
*/
var a []int
var b = []int{1,2,3,4,5}	// array = [1,2,3,4,5] whose descriptor is b

func Slice() {
	fmt.Println("nil slice :", a)
	a := append(a, 100)  // go creates new array = [100] for slice a
	b := append(b, 144)

	fmt.Println(a)
	fmt.Println(b)

	a = b
	// inside a, descriptor of b is stored
	// now both a,b point to same array ---> array of b
	fmt.Println(a)
	fmt.Println(b)

	// change in any one slice refelcts in all who point same array

	s := []byte("golang")
	fmt.Println(s)
}