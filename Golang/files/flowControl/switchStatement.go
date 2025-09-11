package flowControl
import "fmt"

func Switch() {
	d := 7
	switch d {
	case 6, 7:
		fmt.Println("its weekend!")
	case 1,2,3,4,5:
		fmt.Println("Regular working day")
	default:
		fmt.Println("enter within limit!")
	}

	// switch on true
	x := 4
	switch {
	case x <= 2:
		fmt.Println("very less num")
	case x <= 5:
		fmt.Println("num is very close")
	default:
		fmt.Println("try again")
	}
}