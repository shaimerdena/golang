package main

import (
	"fmt"
)

func safeDivide(a, b int) (result int, err error){
	defer func() {
		if r := recover(); r != nil{
			err = fmt.Errorf("recovered: %w", r)
		}
	}()
	return a/b, nil
}

// how a panic travels
func level3(){
	defer fmt.Println("defer in level3")
	panic("something went wrong deep inside")
}
func level2(){
	defer fmt.Println("defer in level2")
	level3()
	fmt.Println("after level3")	// stops at panice - never runs
}
func level1(){
	defer func(){
		if r:=recover(); r!=nil{
			fmt.Println("recovered in level1:", r)
		}
	}()
	level2()
	fmt.Println("after level2")	// stops at panice - never runs
}

func helper(){
	if r:=recover(); r!=nil{
		fmt.Println("recovered:", r)
	}
}
func a(){
	defer func(){
		helper()
	}()
	panic("boom A")
}
func b(){
	defer helper()
	panic("boom B")
}
func c(){
	fmt.Println("c: recover outside defer returned", recover())
}

func main(){
	fmt.Println(safeDivide(10, 2))
	fmt.Println(safeDivide(1, 0))
	
	level1()
	fmt.Println("main continues normally")

	// a()
	b() 
	c()
}