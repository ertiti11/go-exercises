package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	b, error := ioutil.ReadFile("file.txt")
	if error != nil {
		fmt.Print("a error ocurred")
	} else {
		fmt.Print(string(b))
	}
}
