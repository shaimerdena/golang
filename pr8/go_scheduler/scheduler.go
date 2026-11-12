package main

import (
	"fmt"
	"runtime"
	"runtime/pprof"
	"sync"
	"time"
	"syscall"
)

func busy(duration time.Duration){
	start := time.Now()
	for {
		if time.Since(start) > duration{
			break
		}
	}
}

func blockThread() {
    timeout := syscall.NsecToTimeval(int64(200 * time.Millisecond))
    syscall.Select(0, nil, nil, nil, &timeout)
}

func main(){
	// g - goroutine - your function + its stack + its state
	// m - machine - a real OS thread that executes code
	// p - processor - a virtual CPU that schedules goroutines to run on an machine
	// GOMAXPROCS - the number of p

	fmt.Println(runtime.GOMAXPROCS(0)) //current value of GOMAXPROCS
	fmt.Println(runtime.NumCPU()) //number of CPU cores
	runtime.GOMAXPROCS(2) //set GOMAXPROCS to 2
	// by default, GOMAXPROCS=runtime.NumCPU() - the number of CPU cores

	// 8 goroutines, each busy for 2 seconds, on 2 cores
	var wg sync.WaitGroup
	for i:=0; i<8; i++{
		wg.Add(1)
		go func(){
			defer wg.Done()
			busy(2 * time.Second)
		}()
	}
	wg.Wait()
	// GODEBUG=schedtrace=500 go run scheduler.go - this will print the scheduler trace every 500ms

	// preemption - the ability of the scheduler to interrupt 
	// a running goroutine and schedule another one
	go func(){
		for {
		}
	}()
	// time.Sleep(10 * time.Millisecond)
	fmt.Println("main")
	// GODEBUG=acyncpreemptoff=1 go run scheduler.go - this will disable preemption


	// run("100 x time.Sleep", 100, func() {
	// 	time.Sleep(200 * time.Millisecond)
	// })
	// run("100 x blocking syscall", 100, func() {
	// 	ts := syscall.NsecToTimespec(int64(200 * time.Millisecond))
	// 	syscall.Nanosleep(&ts, nil)
	// })
	// Q: Both versions wait 200 ms. Why does time.Sleep need no new
	// threads, while the system call needs almost one thread per
	// goroutine? Use the words G, M and P
	// A: With a blocking system call, the goroutine G and its M are blocked. 
	// The P can be handed to another M, allowing other goroutines to continue running.

	before := pprof.Lookup("threadcreate").Count()

	wg = sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			time.Sleep(200 * time.Millisecond)
		}()
	}
	wg.Wait()

	after := pprof.Lookup("threadcreate").Count()

	fmt.Println("Sleep - new threads:", after-before)

	before = pprof.Lookup("threadcreate").Count()

	wg = sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			blockThread()
		}()
	}
	wg.Wait()
	after = pprof.Lookup("threadcreate").Count()

	fmt.Println("Nanosleep - new threads:", after-before)
}