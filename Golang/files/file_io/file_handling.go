package file_io

import (
	// "fmt"
	"fmt"
	"os"
	"io"
)

func FileIO() {
	folder := [3]string{"file_io/file1.txt", "file_io/file2.txt", "file_io/file3.txt"}

	for _, filePath := range folder {
		file, err1 := os.Open(filePath)

		if err1 != nil {
			fmt.Fprintln(os.Stderr, err1)		// os.Stderr  - current open file pointing at standard error
			continue
		}
		data, err2 := io.Copy(os.Stdout, file)
		
		if err2 == nil {
			fmt.Println(data)
		} else {
			fmt.Println(err2)
		}
		file.Close()
	}
}