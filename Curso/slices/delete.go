package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(x)
	x = append(x[:2], x[5:]...)
	fmt.Println(x)

}
