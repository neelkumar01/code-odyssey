package datatypes
import "fmt"

const (
	a = 24
	s = "go language"
	l = len(s)
)

var k int

func Const() {
	fmt.Println(a)
	fmt.Println(s)
	fmt.Println(l)
	fmt.Println(k)
}