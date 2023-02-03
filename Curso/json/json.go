package main

import (
	"encoding/json"
	"fmt"
)

type persona struct {
	Nombre   string
	Apellido string
	Edad     int
}

func main() {
	p1 := persona{
		Nombre:   "adrian",
		Apellido: "prieto",
		Edad:     20,
	}

	p2 := persona{
		Nombre:   "daniel",
		Apellido: "prieto",
		Edad:     22,
	}

	personas := []persona{p1, p2}

	


	bs, err := json.Marshal(personas)

	if err != nil{
		fmt.Println("error")
	}

	fmt.Println(string(bs))




}
