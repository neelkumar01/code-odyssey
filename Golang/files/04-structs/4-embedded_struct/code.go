package main

import "fmt"

/*
Embedded structs
- embedding one struct in other
- all variables of embedded struct promoted to top level
*/

type student struct {
	name string
	id int
	class int
}
type participant struct{
	student
	category string
}

func main() {
	student1 := participant{
		student: student{
			name: "neel kumar",
			id: 12,
			class: 9,
		},
		category: "sports",
	}

	fmt.Println(student1.name)
	// name accessed by student1.name not student1.student.name
	
}
