package main

import (
	"errors"
	"fmt"
)

func main() {
	var err error = errors.New("something went wrong")
	fmt.Println(err)
}
