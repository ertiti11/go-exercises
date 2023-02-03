package main


import "fmt"



func comprobar(num1 int, num2 int) (int , int) {
	if (num1>num2){
		return num1-num2, num1/num2
	}else{
		return num1+num2, int(num1*num2)
	}
}



func main()  {
	
	var num1 int
	var num2 int


	fmt.Println("introduce un numero")
	fmt.Scan(&num1)
	fmt.Println("introduce un numero")
	fmt.Scan(&num2)

	fmt.Println(comprobar(num1,num2))

}