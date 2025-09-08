package datatypes
import (
	"fmt"
)
/*
%T - type
%v - give value
*/
var objects int
var name string
var (
	n1 = 2
	n2 = 10.2
)

func Basic() {
	a := 10
	b := 2.4

	fmt.Printf("a => type: %T  value: %v \n", a, a)
	fmt.Printf("b => type: %T value: %[1]v \n", b)
	
	fmt.Println("")
	objects = 2
	name = "neel kumar"
	fmt.Println(objects)
	fmt.Println(name)
}

