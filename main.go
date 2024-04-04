package main

import "fmt"

func main() {
	fmt.Println(" Go lang if else")

	loginCount := 23
	var result string

	if loginCount <10 {

		result = "Regular user"
		
	}else if loginCount >10{
		result = "Watch out"
	}else{
		result = "Exacitiry 10 login count"
	}

	fmt.Println(result)


	if 9%2 ==0 {
		fmt.Println("The number is even")
	}else{
		fmt.Println("the number is odd")
	}

	if num :=3; num<10{
		fmt.Println("the number is less than 10")
	}else{
		fmt.Println("num is not lessthan 10")
	}
}