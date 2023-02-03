package main

import (
	"fmt"
	"os"
)

func main() {
	_, err := os.Stat("file.txt")
	if err == nil {
		fmt.Print("File exist\n")
	} else {
		fmt.Print("File not exist")
	}
}
