package main

import "fmt"

/*
Interface analogy

- interface is like a 2-pin socket
- struct is like a plug
- if the plug has 2-pins, it can fit in socket
- similarly, if struct has required methods, it will satisfy the interface

*/

type speaker interface{
	speak()
}
type listener interface{
	listen()
}

type person struct{
	name string
}
func (she person) speak() {
	fmt.Println(she.name, " is speaking")
}
func (she person) listen() {
	fmt.Println(she.name, " is listening")
}

func main() {
	girl := person{"radha"}
	var s speaker = girl
	var l listener = girl

	s.speak()
	l.listen()
}