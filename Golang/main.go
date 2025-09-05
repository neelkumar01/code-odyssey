package main
import (
	"fmt"
	"golang/datatypes"
	"golang/hello_world"
	"golang/strings"
	"golang/arrays"
	"golang/slice"
	"golang/maps"
)

var x = [...]int{1,2,3}
var y = []int{0,0,0}

func main() {
	fmt.Println("This is main function")
	hello_world.Greet()

	datatypes.Basic()
	datatypes.Conversion()
	datatypes.Error()
	datatypes.Const()

	strings.Concept()
	strings.Structure()
	strings.StringFunctions()
	strings.Argument()

	arrays.Array()

	slice.Slice()

	z := slice.Example(x,y)
	fmt.Println(x,y,z)

	maps.Map()
	maps.Literal()
	maps.Index()
}
