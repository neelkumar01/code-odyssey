package gofunctions
import (
	"fmt"
	"os"
)
/*
defer to,
- close file we opened
- close socket / http request we made
- things got saved
etc...

defer  :  takes function call
*/

func Defer() {
	file, err := os.Open("gofunctions/myFile.txt")

	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()		// this is guaranted to run at function exit
}