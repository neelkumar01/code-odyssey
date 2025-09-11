package flowControl
import "fmt"

func Break() {
	for t:=1; t <= 7; t++ {
		fmt.Printf("connected....  t=%v \n",t)
		if t == 5 {
			break
		}
	}
}