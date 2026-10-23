package main

import (
	"cmp"
	"fmt"
	"strings"
)

type Number interface {
	~int | ~int64 | ~float64
}
func Sum[T Number](nums []T) T {
	var s T		//the zero value of T
	for _, n := range nums {
		s += n
	}
	return s
}

func Max[T cmp.Ordered](a, b T) T{
	if a > b{
		return a
	} 
	return b
}

func Contains[T comparable](items []T, target T)bool{
	for _, v := range items{
		if v == target{
			return true
		}
	}
	return false
}

type Celsius float64
type Number2 interface{
	int | float64	//Celsius doesn't satisfy Number without tilde
}


func (c Celsius) String() string{
	return fmt.Sprintf("%.1f°C", c)
}

func JoinAll[T fmt.Stringer](items []T, sep string) string{
	parts := make([]string, 0, len(items))
	for _, it := range items{
		parts = append(parts, it.String())
	}
	return strings.Join(parts, sep)
}


// // 1.fails to compile: not all types are comparable
// func Contains[T any](items []T, target T) bool {
// 	for _, v := range items {
// 		if v == target { return true }
// 	}
// 	return false
// }

// Solution: comparable instead of any

// // 2. fails to compile: comparable doesn't support '>'
// func Max[T comparable](a, b T) T {
// 	if a > b { return a }
// 	return b
// }

// Solution: cmp.Ordered instead of comparable

// // 3. fails to compile: this kind of interface is used
// // for type constraint, not for variable type
// type Number interface{ ~int | ~float64 }
// var n Number = 5

// Solution: use the interface as type constraint

// // 4. fails to compile: the function doesn't know which
// // type of data to return
// func Zero[T any]() T { var z T; return z }
// x := Zero()

// Solution: specify the type: e.g. x := Zero[int]()


// generic types
type Stack[T any] struct { items []T}
func (s *Stack[T]) Push(v T) { s.items = append(s.items, v)}	
func (s *Stack[T]) Pop() (T, bool){
	var zero T
	if len(s.items) == 0{
		return zero, false
	}
	v := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return v, true
}

type Queue[T any] struct { items []T}
func (q *Queue[T]) Push(v T) { q.items = append(q.items, v)}
func (q *Queue[T]) Pop() (T, bool){
	var zero T
	if len(q.items) == 0 {
		return zero, false
	}
	v := q.items[0]
	q.items = q.items[1:len(q.items)]
	return v, true
}

// func (b Box[T]) Map[U any](f func(T) U) Box[U] 
// syntax error: method must have no type parameters

// correct version:
// func Map[T, U any](items []T, f func(T) U) []U


// tasks
// 1. Sum[T Number](nums []T) T and Average[T Number](nums []T) (float64, bool)
// — false for an empty slice. already did.
func Average[T Number](nums []T) (float64, bool){
	if(len(nums) == 0){
		return 0, false
	}
	var s T = Sum(nums)
	return float64(s)/float64(len(nums)), true
}
// 2. MaxOf[T cmp.Ordered](items []T) (T, bool) — test it with int , string and
// an empty slice.
func MaxOf[T cmp.Ordered](items []T) (T, bool){
	var max T
	if(len(items) == 0){
		return max, false
	}
	for _, v := range items{
		if(v > max){
			max = v
		}
	}
	return max, true
}
// 3.IndexOf[T comparable](items []T, target T) int — -1 if not found.
func IndexOf[T comparable](items []T, target T) int {
	for i, v := range items{
		if(v == target){
			return i
		}
	}
	return -1
}
// 4. Filter[T any](items []T, keep func(T) bool) []T — keep the scores >= 75 .
func Filter[T any](items []T, keep func(T) bool) []T {
	res := make([]T, 0)
	for _, v := range items{
		if keep(v) {
			res = append(res, v)
		}
	}
	return res
}
// 5. Remove the ~ from Number and read the compiler error. Put it back.



func main(){
	fmt.Println(Sum([]int{1, 2, 3})) // 6 — T inferred as int
	fmt.Println(Sum([]float64{1.5, 2.5})) // 4 — T inferred as float64
	fmt.Println(Sum[int64]([]int64{10, 20})) // 30 — T written explicitly

	// Constraint			Allowed types									What you can do with T
	// any					every type										assign, pass, return - that's all
	// comparable			types that support == (no slices, maps, funcs)	= , !=, use as a map key
	// cmp. Ordered(Go1.21)	integers, floats, strings (and ~ of them)		<, >, <=, >=, ==
	// int | float64		exactly int and float64							operations that all of them support(+，<, ...）
	// ~int | ~float64		int, float64 and every type built on them		the same operations
	// fmt.Stringer			types with a String() string method				call .String()

	fmt.Println(Max(3, 4), Max("apple", "banana"))
	fmt.Println(Contains([]string{"go", "rust"}, "go"))

	fmt.Println(JoinAll([]Celsius{20.5, 22, 19.8}, ", "))

	var s Stack[string]
	s.Push("a")
	s.Push("b")
	fmt.Println(s.Pop())
	
	nums := &Stack[int]{}
	nums.Push(10)
	fmt.Println(nums.Pop())

	var q Queue[float64]
	q.Push(1.5)
	q.Push(10.4)
	fmt.Println(q.Pop())

	// generics in standard library
	// slices.Contains([]string{"jo", "just"}, "jo")
	// slices.Max([]int{3,4,5,10})
	// slices.Sort(nums)
	// maps.Keys(m)
	// max(2,41,5)

	fmt.Println(Average([]int{2,3,4,5,7}))
	fmt.Println(Average([]int64{}))
	fmt.Println(MaxOf([]int{8,4,1,123}))
	fmt.Println(IndexOf([]string{"jo", "go", "golang"}, "o"))
	scores := []int{12,34,45,65,76,87,98}
	passed := Filter(scores, func(score int) bool {
		return score >= 75
	})
	fmt.Println(passed)

	var a, b any = 1, 1
	fmt.Println(a == b)	//true
	var c, d any = 1, int64(1)
	fmt.Println(c == d)	//false
	// var e, f any = []int{1}, []int{1}
	// fmt.Println(e == f)	//won't compile, bc slices are uncomparable
}