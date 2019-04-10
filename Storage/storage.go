package storage

import (
	"fmt"
	"os"
)

type test struct {
	X int
	Y int
}

//write file method.
func writeFile() {
	file, err := os.Create("test")
	defer file.Close()

	if err != nil {
		fmt.Println("Error!")
	}

	pi := []byte{115, 111, 109, 101, 10}
	file.Write(pi)
}
