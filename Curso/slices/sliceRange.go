package main

import "fmt"

func main() {
	x := []int{5, 2, 3, 5, 6, 7}

	fmt.Println(x)
	/*esto sirve para indicar que rango queremos mostrar
	por ejemplo, este caso imprimiria el valor desde la pos 1 al 3*/
	fmt.Println(x[1:3])

}
