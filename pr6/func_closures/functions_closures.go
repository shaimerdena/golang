package main

import (
	"fmt"
	"strings"
)

// functions as values
func apply(nums []int, f func(int) int) []int{
	out := make([]int, len(nums))
	for i, n := range nums{
		out[i] = f(n)
	}
	return out
}
func double(n int) int { return 2 * n}

// func type with a name
type Op func(a, b float64) float64

// multiple results
func dvmod(a, b int) (q, r int) {
	q = a/b
	r = a%b
	return //"naked" return: returns q and r
}

// variadic: nums is a []int
func sum(nums ...int) int{
	total := 0
	for _, v := range nums{
		total += v
	}
	return total
}


// closures
func counter() func() int {
	count := 0
	return func() int{
		count++
		return count
	}
}

// closure that builds function
func multiplier(k int) func (int) int{
	return func(n int) int { return n * k}
}


// tasks
// 1. Write makeAccumulator() func(int) int . Every call adds the argument to a
// running total and returns the total: acc(10), acc(5), acc(-3) → 10 15 12.
func makeAccumulator() func(int) int{
	total := 0
	return func(i int) int {
		total += i
		return total
	}
}
// 2. Write memoize(f func(int) int) func(int) int . It returns a function that
// remembers results in a map (Week 3). Prove with a counter that the slow
// function is called only once per input.
func memoize(f func(int) int) func(int) int{
	m := make(map[int]int, 0)
	return func(i int) int {
		if val, ok := m[i]; ok{
			return val
		}
		m[i] = f(i)
		return m[i]	
	}
}


func main(){
	fmt.Println(apply([]int{1,2,3}, double))
	//anon function as a value
	fmt.Println(apply([]int{1,2,3}, func(n int) int {return n*n})) 

	// built-in function as a value
	var f func(string) string = strings.ToUpper
	fmt.Println(f("go"))

	ops := map[string]Op{
	"+": func(a, b float64) float64 {return a+b},
	"*": func(a, b float64) float64 {return a*b},
	}
	fmt.Println(ops["+"](2,3), ops["*"](2,3))

	q, r := dvmod(17, 5)
	fmt.Println(q,r)

	fmt.Println(sum(), sum(1, 2, 3), sum([]int{1,2,3}...))

	c1 := counter()
	fmt.Println(c1(), c1(), c1())  //1 2 3
	c2 := counter()
	fmt.Println(c2(), c1())	//1 4

	// Predict both lines. count is a local variable of counter() , and
	// counter() has already returned. Why does count still exist?
	// bc the internal function remembers the outside variable

	x := 10
	show := func() { fmt.Println("x =", x)}
	x = 20
	show()	//20

	triple := multiplier(3)
	fmt.Println(triple(5)) //15

	// loops
	var funcs []func()
	for i:=0; i<3; i++ {
		funcs = append(funcs, func() {fmt.Println(i, " ")})
	}
	for _, f := range funcs{
		f()
	}


	// tasks
	m := makeAccumulator()
	fmt.Println(m(10), m(5), m(-3))

	memo := memoize(double)
	fmt.Println(memo(5))
}