package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

func say(msg string){
	for i:=0; i<3; i++ {
		fmt.Println(msg, i)
		time.Sleep(100 * time.Millisecond)
	}
}

func main(){
	go say("Hello from goroutine")
	say("Hello from main")

	// Q: Where are the goroutine lines? 
	// A: the main function finishes before the goroutine has a chance to run, 
	// so the program exits before the goroutine can print anything.
	
	// Q: Add time.Sleep(100 *
	// time.Millisecond) inside the loop (as in the Tour of Go) and run it
	// several times. Why can the output be diﬀerent every time, and
	// why is goroutine 2 sometimes missing?
	// A: The output can be different every time because goroutines run concurrently, and their execution order is not guaranteed. 
	// The scheduler may choose to run one goroutine before another, 
	// leading to different interleavings of the output. 
	// Goroutine 2 may sometimes be missing because the main function may finish executing 
	// before goroutine 2 has a chance to run, causing the program to exit before it can print its output.

	// waiting for goroutines to finish using WaitGroup
	var wg sync.WaitGroup
	count := 0
	var counter atomic.Int64
	for i:=1; i<=5; i++{
		wg.Add(1) // 1. count the goroutine before it starts
		go func(){
			// wg.Add(1) 	// wrong: Wait() may run before this line
			defer wg.Done()	// 2. say that the goroutine is done when it finishes
			fmt.Println("worker", i, "done")
			count++ 		// wrong: 1000 goroutinges change one variable concurrently: data race
			// fmt.Println("worker", i, "done, count:", count)
			counter.Add(1) 	// correct: atomic operations are safe for concurrent use
			// fmt.Println("worker", i, "done, counter:", counter.Load())
		}()
	}
	wg.Wait() // 3. block until all goroutines are done
	fmt.Println("all workers finished")
	fmt.Println("count:", count)
	fmt.Println("counter:", counter.Load())
	// test concurrent code with the race detector: go run -race routine.go | go test -race routine.go

	
	const N=100_000
	stop := make(chan struct{})
	for i:=0; i<N; i++{
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			<- stop // block until the channel is closed
		}(i)
	}
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	close(stop) // close the channel to signal all goroutines to stop
	fmt.Println("alive goroutines:", runtime.NumGoroutine())
	fmt.Println("KB per goroutine:", float64(m.Alloc)/float64(runtime.NumGoroutine())/1024)
	// goroutines are cheap, but they are not free: 
	// they consume memory and CPU time.

	go func() {
		defer func() {
			fmt.Println("goroutine recovered: ", recover()) // recover from panic
		}()
		var m map[string]int
		m["key"] = 42 // panic: assignment to entry in nil map
	}()

	// panic in a goroutine: the program will crash
	// the goroutine panic should be recovered in the goroutine itself, not in the main function
	defer func() { fmt.Println("main recovered: ", recover())}()
	go func() {
		var m map[string]int
		m["key"] = 42 // panic: assignment to entry in nil map
	}()
	time.Sleep(100*time.Millisecond)
}