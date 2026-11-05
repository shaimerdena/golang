package main

import (
	"fmt"
)

type Point struct {
	X, Y int
}

var global *Point

func f1() int { x := 42; return x }	//stack
func f2 () *int { x := 42; return &x }	//heap
func f3 (p *Point) int {return p.X + p.Y }	//stack
func f4 () { p := Point{1, 2}; fmt.Println(p) }	//stack -> heap
func f5() []int { s := make ([]int, 10); return s }	//heap
func f6() int { s := make([]int, 10); return len(s) }	//stack
func f7(n int) int { s := make([]int, n); return len(s) }	//stack
func f8() func() int {x := 0; return func() int { x++; return x } }	// heap
func f9 (p *Point) { global = p }	//heap
func f10() byte { var big [200_000]byte; return big[0] } //heap

func main(){
	// go build -gcflags '-m' . - short report
	// go build -gcflags '-m -m' . - detailed report
	// go build -gcflags '-m -l' . - disable inlining: easier to read
	// go build -gcflags '-m -l -m' . - detailed report, no inlining

	p := Point{3, 4}
	f1()
	f2()
	f3(&p)
	f4()
	f5()
	f6()
	f7(10)
	f8()
	f9(&p)
	f10()
}