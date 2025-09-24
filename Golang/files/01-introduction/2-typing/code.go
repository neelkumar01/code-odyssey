package main
import "fmt"

func main() {
	var user string = "neel kumar"
	var password string = "123"  
	// error if int instead of string in password
	fmt.Println(user + " : " + password)
}