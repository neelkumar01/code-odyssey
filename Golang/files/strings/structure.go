package strings
import "fmt"

func Structure() {
	s := "I am neel"
	fmt.Println(s)
	
	length := len(s)
	name := s[5:]
	s += " kumar"

	fmt.Println(s)
	fmt.Println(length)
	fmt.Println(name)
}