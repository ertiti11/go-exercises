package main

import "fmt"

func main() {
	x := make([]int, 10, 12)

	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	// no se puede por que sobrepasa la capacidad pero con append si
	//x[15] = 123

	x = append(x, 123, 124, 12, 31, 41, 23, 1, 2, 2, 34)
	fmt.Println(x)
}
