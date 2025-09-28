package main

import "fmt"

/*
Anonymous structs
- when uou create then without name
- you only need them once
- don't need them in future
*/

func main() {
	lion := struct {
		name string
		color string
		area string
	} {
		name: "simba",
		color: "yellow",
		area: "savana",
	}

	fmt.Println(lion)
	
}
