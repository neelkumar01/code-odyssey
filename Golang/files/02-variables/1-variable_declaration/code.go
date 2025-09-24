package main
import "fmt"

 func main() {
	var smsLimit int
	var costPerSms float64
	var hasPermission bool
	var user string

	fmt.Printf(
		"%v, %f, %v, %q",
		smsLimit,
		costPerSms,
		hasPermission,
		user,
	)
 }