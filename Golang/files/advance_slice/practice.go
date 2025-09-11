package advance_slice
import (
	"fmt"
)

func LenCap() {
	a := [3]int{1, 2, 3}		// [1 ,2, 3]
	b := a[0:1]						// [1]

	fmt.Println(a)
	fmt.Println(b)

	c := b[0:2]					// [1, 2] how?  --->  takes extra elements from "a"
	fmt.Println(c)		

	fmt.Println("len b :",len(b))		// 1
	fmt.Println("cap b :",cap(b))		// 3


}