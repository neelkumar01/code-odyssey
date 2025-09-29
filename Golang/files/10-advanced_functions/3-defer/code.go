package main

import (
	"fmt"
	"io"
	"os"
)

func readFile(file string) (data string, err error) {
	src, err1 := os.Open(file)
	if err1 != nil {
		return
	}
	defer src.Close()

	content, err2 := io.ReadAll(src)
	if err2 != nil {
		return
	}
	data = string(content)
	return data, nil
}

func main() {
	// defer - runs a function just before enclosing function returns
	
	fmt.Println(readFile("03-functions/8-early_returns/code.go"))
}