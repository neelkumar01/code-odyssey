package flowControl
import "fmt"

func Traversing() {
	myArray := [...]int{10,20,30}

	for index := range myArray {
		fmt.Println("index:", index)
	}

	for i, val := range myArray {
		fmt.Printf("index: %v  val: %v \n", i, val)
	}
}