package main

import "fmt"

func main() {
	fmt.Println(" Go lang if else")

	loginCount := 9
	var result string

	if loginCount ==10 {

		result = "Regular user"
		
	}else if loginCount >10{
		result = "Watch out"
	}else{
		result = "Exacitiry 10 login count"
	}

	fmt.Println(result)
}