package main

import (
	"fmt"
	"time"
)

func run(name string, ch chan int){
	start := time.Now()
	since := func() time.Duration { return time.Since(start).Round(100 * time.Millisecond)}
	done := make(chan struct{})
	go func() {
		for i := 1; i <= 3; i++ {
			ch <- i
			fmt.Printf("%s: sent %d at %v\n", name, i, since())
		}
		close(done)
	}()
	time.Sleep(300 * time.Millisecond)
	for i:=0; i < 3; i++ {
		<-ch
		time.Sleep(100 * time.Millisecond)
	}
	<- done
}

func main(){
	// unbuf := make(chan int) 	// cap 0: a hand-over
	// buf := make(chan int, 2) 	// cap 3: a small queue
	
	// 						Unbuffered									Buffered (cap = n)
	// Send blocks…			Until a receiver takes the value			Until a free buffer slot is available
	// Receive blocks…		Until a sender sends a value				Until a value becomes available in the buffer
	// Guarantee			The value was received when send returns	The value is in the queue when send returns
	// Think of it as		Handing a book to a person					Putting a letter in a mailbox

	run("unbuffered:", make(chan int))
	run("buffered(3):", make(chan int, 3))

	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	fmt.Println("len: ", len(ch), "cap:", cap(ch))
	close(ch)
	v, ok := <- ch
	fmt.Println(v, ok)
	v, ok = <-ch
	fmt.Println(v, ok)
	v, ok = <-ch
	fmt.Println(v, ok)

	// A buﬀer is not a fix for a deadlock
	// If a program deadlocks with an unbuﬀered channel, adding a buﬀer often
	// only hides the bug until there are more values than slots. Choose the
	// buﬀer size for a reason: "exactly one result" ( cap 1 , Week 8 leak fix),
	// "N workers" ( cap N ), "smooth out short bursts". Don't guess 1000 "just in case".

	// tasks
	// 1. Run run(...) with capacity 0, 1, 2, 3 and 10. Fill a table with the times
	// of the three sends. Explain each row.
	run("buffered(0):", make(chan int, 0))
	run("buffered(1):", make(chan int, 1))
	run("buffered(2):", make(chan int, 2))
	run("buffered(3):", make(chan int, 3))
	run("buffered(10):", make(chan int, 10))
	// | cap | send 1 | send 2 | send 3 |
	// | 0   | 300ms  | 400ms  | 500ms  |	- No buffering
	// | 1   | 0ms    | 300ms  | 400ms  |	- One value can be buffered, other two must wait for a receiver
	// | 2   | 0ms    | 0ms    | 300ms  |   - Two values can be buffered, the third must wait for a receiver
	// | 3   | 0ms    | 0ms    |  0ms   |   - All three values can be buffered
	// | 10  | 0ms    | 0ms    |  0ms   |	- All three values can be buffered, and there is still space for more

	
	// 	2. What does this print, and why doesn't it deadlock? Then change 2 to
	// 1 and explain the new result.
	ch_str := make(chan string, 1)
	ch_str <- "a"
	ch_str <- "b"
	fmt.Println(<-ch_str, <-ch_str)
	// A: prints: a b, when buffer equals to 2
	// B: prints: a, when buffer equals to 1, 
	// because the second send blocks until the first value is received, 
	// and since there is no receiver for the second value, it will not be sent 
	// and will cause a deadlock if we try to receive it.
}