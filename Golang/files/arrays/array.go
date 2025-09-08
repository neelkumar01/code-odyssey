package arrays
import "fmt"

var a  [3]int	// [ 0, 0, 0]
/*
-  in above code, we create array
-  size  =  3
-  contains 3 places to store value of type = int
-  default values put to 0
*/
var b = [3]int{1,2,3}	// assigning values in array

var c = [...]int{1,2,3,4,5,6}	// making compiler decide how many values given to array

func Array() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	
}
