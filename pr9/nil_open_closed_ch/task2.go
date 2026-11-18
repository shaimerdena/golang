package main

import "fmt"

// tasks
// 2. Write a tiny program for each cell of the table above (12 programs, or
// one program with 12 functions) and check the result. Which cells end with
// panic , which with fatal error: ... deadlock ?

// nil channel
func nilSend(){
	var ch chan int
	ch <- 42
}
func nilReceive(){
	var ch chan int
	fmt.Println(<-ch)
}
func nilClose(){
	var ch chan int
	close(ch)
}
func nilLenCap(){
	var ch chan int
	fmt.Println(len(ch), cap(ch))
}


// open channel
func openSend(){
	ch := make(chan int)
	ch <- 42
}
func openReceive(){
	ch := make(chan int)
	fmt.Println(<-ch)
}
func openClose(){
	ch := make(chan int)
	close(ch)
}
func openLenCap(){
	ch := make(chan int)
	fmt.Println(len(ch), cap(ch))
}


// closed channel
func closedSend(){
	ch := make(chan int)
	close(ch)
	ch <- 42
}
func closedReceive(){
	ch := make(chan int)
	close(ch)
	fmt.Println(<-ch)
}
func closedClose(){
	ch := make(chan int)
	close(ch)
	close(ch)
}
func closedLenCap(){
	ch := make(chan int)
	close(ch)
	fmt.Println(len(ch), cap(ch))
}



func main(){
	// nilSend()		// fatal error: all goroutines are asleep - deadlock!
	// nilReceive()		// fatal error: all goroutines are asleep - deadlock!
	// nilClose()		// panic: close of nil channel
	// nilLenCap()		// 0 0
	// openSend() 		// fatal error: all goroutines are asleep - deadlock! (because no one is receiving)
	// openReceive()	// fatal error: all goroutines are asleep - deadlock! (because no one is sending)
	// openClose()		// ok
	// openLenCap()		// 0 0
	// closedSend()		// panic: send on closed channel
	// closedReceive()	// panic: send on closed channel
	// closedClose()	// panic: close of closed channel
	// closedLenCap()	// 0 0
}