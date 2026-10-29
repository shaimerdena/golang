package main

import (
	"fmt"
)

func mustPositive(n int) int{
	if n <= 0 {
		panic(fmt.Sprintf("mustPositive: got %d", n))
	}
	return n
}

// tasks
// Task 6: Safe task runner
// Write SafeRun(tasks []Task) []error , where type Task struct { Name string;
// Run func() error }.
// Every task runs, even if the previous one panicked.
// A returned error becomes task <name>: <error> (wrapped with %w ).
// A panic becomes task <name>: panic: <value> .
// Test it with tasks that: succeed; index out of range; write to a nil map;
// call panic("config is missing") ; return a sentinel error. Then check with
// errors.Is that the sentinel can still be found inside errors.Join(errs...).
type Task struct {
	Name string
	Run  func() error
}


func main(){
	// s := []int{1,2,3}
	// fmt.Println(s[5])	// panic: runtime error: index out of range [5] with length 3

	// var m map[string]int
	// m["x"] = 1	// panic: assignment to entry in nil map

	// var p *Student
	// fmt.Println(p.Name)	// panic: runtime error: invalid memory address or nil pointer dereference

	// a, b := 1, 0
	// fmt.Println(a / b) // panic: runtime error: integer divide by zero

	n := 2
	fmt.Println(mustPositive(n))
}