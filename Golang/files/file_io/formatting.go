package file_io
import "fmt"
/*
standard input
standard output
standard error
*/

func Formatting() {
	a, b := 12, 247
	c, d := 1.2, 2.47

	fmt.Println("\n\n Formatting...")

	fmt.Printf("a: %d   b: %d \n", a,b)
	fmt.Printf("a: %#x  b: %#x \n", a,b)	// base16
	fmt.Printf("c: %f  d:%.4f \n", c, d)

	fmt.Println()

	fmt.Printf("a : |%10d|  b : |%10d|", a, b)
	fmt.Printf("a : |%010d|  b : |%010d|", a, b)
	fmt.Printf("a : |%-10d|  b : |%-10d|", a, b)  // left justify

	fmt.Println()
	s := []int{1,2,3}

	fmt.Printf("%T\n", s)
	fmt.Printf("%v\n", s)
	fmt.Printf("%#v\n", s)

	fmt.Println()
	arr := [5]int{12, 24, 36, 48, 60}
	fmt.Printf("%T\n", arr)		// type
	fmt.Printf("%v\n", arr)		
	fmt.Printf("%#v\n", arr)
}