package main

import "fmt"

func add(a int, b int) int {
	return a + b
}
func concat(s1 string, s2 string) string {
	return s1 + s2
}
func main() {
	v1 := add(2,7)
	fmt.Println(v1)

	v2 := concat("hello", "world")
	fmt.Println(v2)
}
