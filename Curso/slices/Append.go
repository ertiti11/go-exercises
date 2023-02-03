package main

import "fmt"

func main() {
	x := []int{5, 2, 3, 5, 6, 7}
	y := []int{3, 2, 1, 1, 3, 3}
	fmt.Println(x)
	x = append(x, 99, 67, 564, 345)
	x = append(x, y...)
	fmt.Println(x)
}
