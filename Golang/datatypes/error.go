package datatypes
import (
	"fmt"
	"os"
)

func Error() {
	file, myError := os.Open("data.txt")
	// checking for error....
	if myError != nil {
		fmt.Println("Error: ", myError)
	} else {
		fmt.Println(file.Name())
		file.Close()
	}
}
