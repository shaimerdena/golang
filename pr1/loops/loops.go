package main

import "fmt"

func main(){
	// for loop
	for i:=0; i<10; i++{
		fmt.Println(i)
	}
	// for works like a while loop
	sum := 100
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}