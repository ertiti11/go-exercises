package main

import "fmt"

func main() {
	for i := 1; i <= 90; i++ {
		fmt.Print(i)
		fmt.Printf("\t%#U\n", i)
	}
}
