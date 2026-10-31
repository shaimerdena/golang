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
func SafeRun(tasks []Task) []error {
	var errs []error
	for _, task := range tasks {
		func(t Task) {
			defer func() {
				if r := recover(); r != nil {
					errs = append(errs, fmt.Errorf("task %s: panic: %v", t.Name, r))
				}
			}()
			if err := t.Run(); err != nil {
				errs = append(errs, fmt.Errorf("task %s: %w", t.Name, err))
			}
		}(task)
	}
	return errs
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


	tasks := []Task{
		{
			Name: "task1",
			Run: func() error {
				fmt.Println("task1: success")
				return nil
			},
		},
		{
			Name: "task2",
			Run: func() error {
				fmt.Println("task2: panic")
				panic("config is missing")
			},
		},
		{
			Name: "task3",
			Run: func() error {
				fmt.Println("task3: index out of range")
				s := []int{1, 2, 3}
				fmt.Println(s[5]) // panic
				return nil
			},
		},
		{
			Name: "task4",
			Run: func() error {
				fmt.Println("task4: write to nil map")
				var m map[string]int
				m["x"] = 1 // panic
				return nil
			},
		},
		{
			Name: "task5",
			Run: func() error {
				fmt.Println("task5: sentinel error")
				return fmt.Errorf("sentinel error")
			},
		},
	}

	errs := SafeRun(tasks)
	for _, err := range errs {
		fmt.Println(err)
	}

}