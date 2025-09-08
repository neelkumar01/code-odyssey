// writing a simple search and replace program

package strings

import (
	"fmt"
	"os"
)

func Argument() {
	if len(os.Args) > 1 {
		currentPath := os.Args[0]
		fmt.Println("\ncurrent path :", currentPath)
		name := os.Args[1]
		fmt.Println("\nname :", name)
	}
	/*
		=>  go run main.go Neel
		for os.Args,
		os.Args[0]  =  file path
		os.Args[1]  =  next argument we give after command in terminal ---> 'Neel'
	*/
}
