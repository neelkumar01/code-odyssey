package main

import (
	"errors"
	"fmt"
)

func main() {
	x, y, errorMsg := calculate(8, 0)
	if errorMsg != nil {
		fmt.Println(errorMsg)
	} else {
		fmt.Println("8*4 = ", x)
		fmt.Println("8/4 = ", y)
	}
}

func calculate(a, b int) (mul, div int, err error) {
	if b == 0 {
		return 0, 0, errors.New("Division by zero is invalid")
	}
	mul = a * b
	div = a / b
	return mul, div, nil
}
