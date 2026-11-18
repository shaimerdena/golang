package main

import (
	"fmt"
	"time"
)

// func producer(out chan<- int) // send-only
// func consumer(in <-chan int) // receive-only
func producer(out chan<- int) {
	out <- 42
}
func consumer(in <-chan int) {
	value := <- in
	fmt.Println(value)
}

// 3. Rewrite the first goroutine example of Week 8 ( say("goroutine") ) so
// that main waits for the goroutine using a channel instead of a WaitGroup.
func say(msg string) chan string {
	ch := make(chan string, 3)
	for i:=0; i<3; i++ {
		ch <- msg
	}
	return ch
}


func main(){ 
	start := time.Now()
	since := func() time.Duration { return time.Since(start).Round(100 * time.Millisecond) }

	ch := make(chan string)
	go func() {
		time.Sleep(500 * time.Millisecond)
		fmt.Println("sender: about to send", since())
		ch <- "ping"
		fmt.Println("sender: about to send", since())
	}()
	fmt.Println("receiver: waiting...", since())
	msg := <- ch
	fmt.Println("receiver: waiting...", msg, "  ", since())
	// sender sleeps for 500ms

	ch1 := make(chan int)
	go producer(ch1)
	consumer(ch1)


	// tasks
	// 1. Run the ping example. Then move time.Sleep(500 * time.Millisecond)
	// from the sender to the receiver (before <-ch ). What changes in the
	// output and why?

	start = time.Now()
	since = func() time.Duration { return time.Since(start).Round(100 * time.Millisecond) }

	ch = make(chan string)
	go func() {
		fmt.Println("sender: about to send", since())
		ch <- "ping"
		fmt.Println("sender: about to send", since())
	}()
	fmt.Println("receiver: waiting...", since())
	time.Sleep(500 * time.Millisecond)
	msg1 := <- ch
	fmt.Println("receiver: waiting...", msg1, "  ", since())
	// receiver sleeps for 500ms


	// | Operation            | Nil channel                       | Open channel                               | Closed channel                                                                  |
	// | -------------------- | --------------------------------- | ------------------------------------------ | ------------------------------------------------------------------------------- |
	// | `ch <- v`            | Blocks forever                    | Blocks until a receiver / free buffer slot | **Panic:** `send on closed channel`                                             |
	// | `<-ch`               | Blocks forever                    | Blocks until a value is available          | Returns remaining values, then the **zero value** and `ok == false` immediately |
	// | `close(ch)`          | **Panic:** `close of nil channel` | OK                                         | **Panic:** `close of closed channel`                                            |
	// | `len(ch)`, `cap(ch)` | `0`                               | Values in the buffer / buffer size         | Values left / buffer size                                                       |

	// task 3
	ch = make(chan string)
	ch = say("goroutine")
	for i:=0; i<3; i++ {
		fmt.Println(<-ch)
	}

	// task 4
	// Q: why a deadlock in main is
	// reported by the runtime, but a goroutine blocked forever on a channel
	// inside a running server is not (Week 8, leaks).
	// A: The runtime can detect deadlocks in the main goroutine because it is the entry point of the program. 
	// If the main goroutine is blocked and there are no other goroutines that can make progress, 
	// the runtime concludes that the program is deadlocked. 
	// In contrast, a goroutine blocked on a channel inside a running server may not be detected 
	// as a deadlock because there could be other active goroutines 
	// that are still able to make progress, 
	// even if one specific goroutine is blocked. 
	// The runtime does not have enough information to determine that 
	// the blocked goroutine is causing a deadlock in the context of the entire program.
}