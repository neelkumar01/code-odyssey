package structs

import "fmt"

// struct contains different data types

type Employee struct {
	name string
	age  int
	id   int
}
type animal struct {
	specie     string
	color      string
	totalCount *count // pointer to be placed
}
type count struct {
	population int
}

func Struct() {
	var e Employee
	e.name = "neel"
	e.age = 18
	e.id = 7
	fmt.Println(e)

	e2 := Employee{
		name: "raam",
		age:  19,
		id:   8,
	}
	fmt.Println(e2)

	e3 := Employee{"krishna", 24, 1}
	fmt.Println(e3)

	// example 2
	a := animal{
		specie:     "lion",
		color:      "yellow",
		totalCount: &count{population: 10},
	}
	fmt.Println(a)

	// example 3
	// using map
	type student struct {
		rollNo  int
		class   int
		section string
	}
	team := map[string]*student{}
	// key type is string
	// in value, we store a pointer that gives us address of student
	team["neel"] = &student{27, 12, "B"}
	team["siddhant"] = &student{45, 12, "D"}

	// example 4
	var album = struct {
		title  string
		artist string
		year   int
	}{
		"personal album",
		"abc singer",
		2025,
	}
	fmt.Println(album)

	var albumPointer = &album
	fmt.Println("album pointer :", albumPointer)

	// example 5
	var menu1 = struct {
		region string
	}{
		"North indian",
	}
	var menu2 = struct {
		region string
	}{
		"South indian",
	}

	menu1 = menu2		// both have type struct so no error !
	fmt.Println(menu1)
	fmt.Println(menu2)

	// example 6  --> doing eg5 with different way
	type catalog1 struct {
		category string
	}
	type catalog2 struct {
		category string
	}
	var c1 = catalog1{"snacks"}
	var c2 = catalog2{"drinks"}

	fmt.Println(c1)
	fmt.Println(c2)

	//   c1 = c2   --->   this will give error as type of both is different
	// type of c1  =  catalog1
	// type of c2  =  catalog2

	//  c1 = c2   --->   possible with type conversion :-
	c1 = catalog1(c2)
	fmt.Println(c1)
	fmt.Println(c2)
}
