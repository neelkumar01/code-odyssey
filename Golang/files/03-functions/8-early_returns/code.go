package main

import "fmt"

func main() {
	id := getId1("neel")
	fmt.Println(id)

	userId := getId2("ram")
	fmt.Println(userId)
}

// 	nested conditions
func getId1(name string) (id int) {
	if name == "neel" {
		id = 1
	} else {
		if name == "ram" {
			id = 2
		} else {
			if name == "krishna" {
				id = 3
			}
		}
	}
	return id
}

//  clean code with early return :
func getId2(name string) (id int) {
	if name == "neel" {
		return 1
	}
	if name == "ram" {
		return 2
	}
	if name == "krishna" {
		return 3
	}
	return 0	// default value
}
