package main

import "fmt"


func main()  {


	var num int

	fmt.Scan(&num)

	if num<10 && num >0{
		fmt.Println("yes")
	}else{
		fmt.Println("no")
	}
	
}