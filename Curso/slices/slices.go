package main

import "fmt"

func main() {
	x := []int{4, 5, 87, 3, 1, 4}

	for indice, valor := range x {
		fmt.Print(indice)
		fmt.Print("\t", valor, "\n")
	}
}
