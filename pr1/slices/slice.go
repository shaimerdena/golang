package main

import (
	"fmt"
)

func main(){
	var slice1 []int
	slice1 = append(slice1, 20, 15, 10, 8)
	fmt.Println(slice1)

	for _, v := range slice1{
		fmt.Print(v, " ")
	}
	fmt.Println()
	for i, _ := range slice1 {
		fmt.Print(i, " ")
	}
	fmt.Println()

	var numbers []int
	numbers = append(numbers, 10, 2, 9, 77)
	slice1 = append(slice1, numbers...)
	fmt.Println(slice1)

	var slice2 []int = slice1[1:4]
	fmt.Println(slice2)

	s := make([]int, 3, 5) 	//make
	fmt.Println(s)      
	fmt.Println(len(s))		//length
	fmt.Println(cap(s))		//capacity

	for a := 0; a < 5; a++ {
		fmt.Println("len -", len(s), "cap -", cap(s))
		s = append(s, a)
	}

	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}
	board[0][0] = "X"
}