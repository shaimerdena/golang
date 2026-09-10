package main

import (
	"fmt"
	"math"
)

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func adder() func(int) int{		//closure function
	sum := 0
	return func(x int) int{
		sum += x
		return sum
	}
}

func fib() func() int {
	i, j := 0, 1
	return func() int{
		curr := i
		i, j = j, i+j
		return curr
	}
}

func main(){
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(3,4))
	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))

	fmt.Println("Adder: ")
	pos := adder()
	for i:=0; i<10; i++{
		fmt.Print(pos(1), " ")
	}
	fmt.Println()

	fmt.Println("Fibonacci: ")
	f := fib();
	for i := 0; i<10; i++ {
		fmt.Print(f(), " ")
	}
}