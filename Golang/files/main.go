package main

import (
	"fmt"
	"golang/arrays"
	"golang/datatypes"
	"golang/hello_world"
	"golang/maps"
	"golang/nil_Pointers"
	"golang/slice"
	"golang/strings"
	"golang/flowControl"
	"golang/goTypes"
	"golang/file_io"
	"golang/gofunctions"
	"golang/closure"
	"golang/advance_slice"
	"golang/structs"
	"golang/regex"
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

	nil_Pointers.Nil()
	fmt.Println("\n\n ")
	nil_Pointers.ReadFile()

	flowControl.IfElse()
	flowControl.ForLoop()
	flowControl.Traversing()
	flowControl.Break()
	flowControl.Label()
	flowControl.Switch()

	goTypes.StructuralTyping()
	goTypes.NamedTyping()

	file_io.Formatting()
	file_io.FileIO()

	gofunctions.Func()
	gofunctions.Pointer()

	myVal := closure.Counter()
	fmt.Println(myVal())
	fmt.Println(myVal())
	fmt.Println(myVal())

	closure.Example()

	advance_slice.AdvSlice()
	advance_slice.LenCap()

	structs.Struct()
	structs.PassingStruct()
	structs.Tags()

	regex.Regex()
}
