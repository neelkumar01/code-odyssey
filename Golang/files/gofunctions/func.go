package gofunctions
import "fmt"

func modify(arr [3]int) [3]int {
	arr[0] = 0
	return arr
}

func Func() {
	myArr := [3]int{10, 20, 30}
	newArr := modify(myArr)

	fmt.Println("myArr :", myArr)
	fmt.Println("newArr :", newArr)
}