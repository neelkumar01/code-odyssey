package goTypes
import "fmt"

type num int 		// "num" becomes a type like int

func NamedTyping() {
	var num1 num		// type of myNum = num not int
	num2 := 12				// type = int

	num1 = num(num2)

	fmt.Println(num1)
}