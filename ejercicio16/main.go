package main

import "fmt"



func media(nota1 float32,nota2 float32, nota3 float32)(float32){
	medias := (nota1+nota2+nota3)/3
	return medias
}
func main()  {
	var nota1 float32
	var nota2 float32
	var nota3 float32
	
	fmt.Print("introduzca nota 1:")
	fmt.Scanln(&nota1)
	fmt.Print("introduzca nota 2:")
	fmt.Scanln(&nota2)
	fmt.Print("introduzca nota 3:")
	fmt.Scanln(&nota3)
	result := media(nota1,nota2,nota3)
	if (result>6.5){
		fmt.Printf("tu media es de media %v" ,result)
	}else{
		fmt.Println("no has promocionado")
	}
}


