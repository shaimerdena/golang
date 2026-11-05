package main

import (
	"fmt"
)

type User struct{
	Name string
	Age int
}
func newUserValue(name string) User {
	u := User{Name: name, Age: 20}
	return u  // a Copy of u goes back to the caller
}
func newUserPointer(name string) *User{
	u := User{Name: name, Age: 20}
	return &u	// a pointer to u goes back
}
// Q: u is a local variable of newUserPointer. The function's stack frame
// disappears when it returns — so where does &u point to? In C this
// is a famous bug ("dangling pointer"). Why is it safe in Go?
// A: bc u escaped to the heap. 

func depth(n int) int{
	var pad [64]byte	// 64KB of padding to make stack frames large
	pad[n%len(pad)] = byte(n)
	if n == 0 {
		return 0
	}
	return depth(n-1) + 1 + int(pad[(n+1)%len(pad)]) //reads a 0, but compiler doesn't know that
}

func forever(n int) int{
	return forever(n+1) + 1
}

// tasks
// 1. Run depth(1 000 000). Then increase the size of pad to [1024]byte. Does
// it still work? Estimate how much stack memory one call uses in each case.
// (You can see the real frame size with go build -gcflags=-S : look for TEXT
// main.depth(SB), ABIInternal, $NN-8.)
// A: depth(1_000_000) works with [64]byte, but fails with [1024]byte. 
// The stack memory used per call is approximately 64 bytes in the first case and 1024 bytes in the second case.

// 2. Remove the line that reads pad and check the frame size again. What did the compiler do?
// A: The compiler optimized the function and removed the unused variable, 
// resulting in a smaller frame size. 
// The stack memory used per call is significantly reduced 
// when the line that reads pad is removed, 
// as the compiler can eliminate the allocation of the pad array entirely.

// 3. Run forever(0) and read the error. Is it a panic that recover can catch (Week 6)? Try it.
// A: Running forever(0) will result in a stack overflow error.
// recover cannot catch a stack overflow panic in Go. 

// Participation Task: explain on the defence why the goroutine stack needs
// to be copied when it grows, and why that is safe even if some pointers
// point into the stack.
// A: In Go, goroutines start with a small stack size and can grow dynamically as needed.
// During the copy, the pointers that point into the stack are updated to point to the new stack location.

func main(){
	u := newUserValue("Alice")
	fmt.Println(u.Name, u.Age)
	
	u2 := newUserPointer("Bob")
	fmt.Println(u2.Name, u2.Age)

	fmt.Println("depth: ", depth(1_000_000))
	// fmt.Println("forever: ", forever(0))	// stack overflow
}