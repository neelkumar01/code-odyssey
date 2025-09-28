package main

import "fmt"

type User struct{
	name string
	age  int
}
func securityCheck(data interface{}) {
	user, ok := data.(User)	// type assertion
	if ok {
		fmt.Println(user.name, user.age)
	} else {
		fmt.Println(data)
	}
}

func main() {
	u := User{name: "neel", age: 18}
	securityCheck(u)
}


/*
user, ok := data.(User)
‣ In securityCheck(arg) function, if type of arg match to User
then,
‣ ok == true
‣ user becomes arg
*/