package main

import "fmt"




func comprobar(clave1 string, clave2 string)(bool){
	if (clave1 == clave2){
		return true
	}else{
		return false
	}
}


func main()  {

	var clave1 string 
	var clave2 string 


	fmt.Print("introduce la clave: ")
	fmt.Scan(&clave1)
	fmt.Print("introduce la clave de nuevo: ")
	fmt.Scan(&clave2)	
	if (comprobar(clave1, clave2)){
		fmt.Println("son iguales")
	}else{
		fmt.Println("las claves son incorrectas")
	}
}

