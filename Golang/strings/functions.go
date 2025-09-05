package strings
import (
	"fmt"
	"strings"
)

func StringFunctions() {
	name := "Hi, neel"
	fmt.Println(len(name))

	fmt.Println(strings.Contains(name, "neel"))
	fmt.Println(strings.Index(name, "Hi"))

	newName := strings.ToUpper(name)
	fmt.Println(newName)
}