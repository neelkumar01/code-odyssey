package closure

import "fmt"

func Example() {
	s := make([]func(), 4)

	for i := 0; i < 4; i++ {
		i2 := i
		s[i] = func() { fmt.Println(i2) } // this is closure  -  a function alongwith variable i2
	}

	for i := 0; i < 4; i++ {
		s[i]() // calling closure
	}
}
