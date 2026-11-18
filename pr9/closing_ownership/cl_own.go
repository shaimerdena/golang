package main

import (
	"fmt"
	"sync"
)

// generator owns its channel: creates it, writes to it, closes it.
func generator(nums ...int) <- chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, n := range nums{
			out <- n
		}
	}()
	return out
}


// tasks
// 2. Write a pipeline of three stages, each one owning its output channel:
// generator(1..10) → square(in <-chan int) <-chan int → sum(in <-chan int)
// int. Every stage must defer close(out).
func generator_task(nums ...int) <-chan int {
	ch := make(chan int)
	go func(){
		defer close(ch)
		for v := range nums {
			ch <- v
		}
	}()
	return ch
}
func square(in <- chan int) <- chan int{
	ch := make(chan int)
	go func(){
		defer close(ch)
		for v := range in {
			ch <- v*v	//works incorrectly - takes indexes instead of values
		}
	}()
	return ch
}
func filterEven(in <- chan int) <- chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for v := range in {
			if(v % 2 == 0){
				ch <- v
			}
		}
	}()
	return ch
}
func summa(in <- chan int) int {
	sum := 0

	for v := range in {
		sum += v
	}
	
	return sum
}

func main(){
	// Rules of closing
	// 1. Only the sender closes — a receiver can't know whether more values
	// are coming.
	// 2. Close once. Closing a closed channel panics.
	// 3. Closing is optional. Close only when a receiver needs to know that the
	// stream has ended ( range , ok ). An unused channel that is not closed is
	// simply garbage-collected.

	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	close(ch)	// broadcast to all receivers that no more values will be sent
	for v := range ch {	//range ends when channel is closed and empty
		fmt.Println(v)
	}

	// the G that creates a channel, writes to it, and closes it is its owner.
	// A function that owns a channel usually returns it as receive-only
	for v := range generator(1, 2, 3){
		fmt.Println(v, " ") // 1 2 3
	}

	results := make(chan int)
	var wg sync.WaitGroup
	for w := 1; w <= 3; w++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 3; i++ {
				results <- w*10 + i
			}
		}()
	}
	// One writer: it closes. 
	// Many writers: the one who waits for all of them closes. 
	// Readers never close.
	go func() {
		wg.Wait()
		close(results)
	}()
	sum := 0
	for v := range results {
		sum += v
	}
	fmt.Println("sum:", sum)
	
	gch := generator_task(1,2,3,4,5)
	summ := summa(square(filterEven(gch)))
	fmt.Println(summ)
}