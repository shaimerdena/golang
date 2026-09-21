package main

import "fmt"

// a Go program that passes a slice to a function
func double(s []int){
	for i := 0; i <len(s); i++{
		s[i] = s[i] * 2
	}
}

func main(){
	// declare an array of 5 float numbers
	var arr [5]float64
	arr[0] = 5.6
	arr[1] = 4.5
	arr[2] = 8.9
	fmt.Println(arr)

	// declare a simple 2D array and for range loop
	var grid [2][2] int = [2][2]int{
		{1, 4},
		{5, 6},
	}
	for i, _ := range grid{
		for j, _ := range grid{
			fmt.Print(grid[i][j], " ")
		}
		fmt.Println()
	}

	cars := []string{"Ferrari", "Honda", "Ford", "BYD"}
	fmt.Println("cars:", cars, "has old length", len(cars), "and capacity", cap(cars)) //capacity of cars is 4
	cars = append(cars, "Toyota")
	fmt.Println("cars:", cars, "has new length", len(cars), "and capacity", cap(cars)) // What is the len and capacity of cars now? Why?
	// cap is 8

	// modifications done to the slice that affects the underlying array 
	var arr5 [5]int = [5]int{5, 6, 8, 9, 10}
	fmt.Println("arr5:", arr5)
	slice6 := arr5[1:3]
	fmt.Println("slice6:", slice6, cap(slice6))
	slice6[0] = 111
	fmt.Println("arr5 afer modification of slice6:", arr5)
	arr5[1] = 1000
	fmt.Println("slice6 after modification of arr5:", slice6)

	// declare a slice in 4 different ways
	s := []int{5,6,7,8,9}
	var v[]int
	s2 := make([]int, 2)
	s3 := s[1:2]
	// -----
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println(v)

	// a slice that was passed to a function
	arr3 := []int{5,6,4,2,2}
	double(arr3)
	fmt.Println(arr3)





	src := []int{1,2,3,4}
	dst := make([]int, 2)
	copy(dst, src)
	fmt.Println(src)
	fmt.Println(dst)
}