package main

import (
	"fmt"
	"pr2/mathutils"
)

func main(){
	a := 10
	b := 100
	fmt.Println(mathutils.Add(a, b) + mathutils.Mult(a, b))
	fmt.Println(mathutils.PercentOf(a, b))
	fmt.Println(mathutils.Swap(a, b))
}