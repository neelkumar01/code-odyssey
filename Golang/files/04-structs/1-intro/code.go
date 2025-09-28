package main

import "fmt"

/*
Struct
- group of variables of different data types
- assigned under one name
*/

type student struct {
	id int
	name string
	class int
	genger string
}

func main() {
	student1 := student {
		id: 1,
		name: "Neel kumar",
		class: 12,
		genger: "male",
	}
	fmt.Println(student1)
	fmt.Println("name : ", student1.name)

	// modify
	student1.name = "krishna"
	fmt.Println(student1)

	
}
