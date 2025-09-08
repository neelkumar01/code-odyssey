package nil_Pointers

import (
	"fmt"
	"os"
	"bufio"
)

func ReadFile() {
	// opening file
	file, err := os.Open("nil_Pointers/test.txt")

	// error handling
	if err != nil {
		fmt.Println("error opening file :", err)
	}
	defer file.Close()

	words := make(map[string]int)
	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanWords)		// to break input to chunks

/*
scanner.Scan()   -  read next token from input and returns true or false if bot found

*/
	for scanner.Scan() {
		words[scanner.Text()]++
	}

	fmt.Println(len(words), "total words / tokens")
}