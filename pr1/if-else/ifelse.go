package main

import "fmt"

func main(){
	var flag bool
	for flag == false{
		var age int
		fmt.Println("Enter your age: ")
		fmt.Scanln(&age)	
		if age <= 0{
			fmt.Println("Error, try again. ")
		} else if age > 0 && age < 18 {
			fmt.Println("You can't enter the system")
			flag = true
		} else{
			fmt.Println("Welcome to the system!")
			flag = true
		}
	}
}