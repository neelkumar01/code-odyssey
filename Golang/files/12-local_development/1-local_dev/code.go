package main

import "github.com/neelkumar01/testmodule"

func main() {
	/*
	main package
	‣ made to directly execute
	‣ entry point : main() function - package starts executing from here

	‣ package other than main - library package like "fmt", "os", "io"
	‣ package - directory of go codes (all compiled together)

	‣ module - collection of packages

	making go.mod file
	‣ go mod init github.com/neelkumar01/gomodule
	printing go.mod
	‣ cat go.mod
	*/
	testmodule.SayHello("neel")
}
