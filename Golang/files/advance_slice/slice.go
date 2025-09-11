package advance_slice
import "fmt"

func AdvSlice() {
	var s []int		// nil
	e := []int{}	// non nil

	fmt.Printf("%d,  %d,  %T,  %t \n", len(s), cap(s), s, s==nil)

	fmt.Printf("%d,  %d,  %T,  %t \n", len(e), cap(e), e, e==nil)
}

/*
arr  --->  [1, 2, 3, 4, 5]
slice  --->  arr[1:4]  =  [2, 3, 4]

length of slice = 3
capacity = 4  ( from starting index of slice till end of array)
*/