package strings

import (
	"fmt"
)

/*
bit ---> 0 or 1
byte ---> 8bits  =  uint8  =  2^8  combinations can be store
int32  --->  32 bits  =  2^32  combinations  =  rune

UTF-8  :  encoding
	=>  a system that convert any char in pattern of 0 or 1
	=>  has unicode for every char
	=>  unicode  --->  univeral code or a unique code for every char


*/

func Concept() {
	name := "Neel"
	fmt.Printf("\n %T  %[1]v", name)

	fmt.Println([]rune(name))
	// Neel[78 101 101 108]
	fmt.Println([]byte(name))
}
