package main

import "fmt"

type loginError struct{
	name string
	err  string
}

func (e loginError) Error() string {
	return fmt.Sprintf("Account access to %v denied for: %v", e.name, e.err)
}

func login(username, password string) error {
	if username!= "neel" {
		return loginError{name: username, err: "user not found"}
	}
	if password != "1234" {
		return loginError{name: username, err: "incorrect password"}
	}
	return nil
}
func main() {
	err := login("neel", "123")
	if err != nil {
		fmt.Println(err)
	}
}

/*
Error()
‣ error is a special interface in go
‣ contains one method - Error()
‣ using Error(), we can convert our struct to error type
*/