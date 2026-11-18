package main

import (
	"fmt"
	"time"
)

func merge(a, b <- chan int) <- chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for a != nil || b != nil {
			select {
			case v, ok := <- a:
				if !ok {
					a = nil
					continue
				}
				out <- v
			case v, ok := <- b:
				if !ok {
					b = nil 
					continue
				} 
				out <- v
			}
		}
	}() 
	return out
}
// Q: What happens without the a = nil line? We counted it: after a
// was closed, the loop received 3039–4397 useless zero values from
// a while waiting for 3 values from b . Explain using the table from Part 1.
// A: nil blocks the channel, so the select statement will not choose that case anymore. 
// Without setting a = nil, the select statement will continue to try to read from the closed channel a, 
// which will always return the zero value (0 for int) and ok as false. 
// this leads to an infinite loop of receiving zero values from a while waiting for values from b.

func producer(words []string, done <- chan struct{}) <- chan string {
	jobs := make(chan string)
	go func() {
		defer close(jobs)
		for _, word := range words {
			select {
			case jobs <- word:
			case <-done:
				return
			}
		}
	}()
	return jobs
}

// tasks
// 2. Write fetchWithTimeout(timeout) that waits for a slow function with
// select + time.After — without leaking the goroutine (Week 8).
func fetchWithTimeout(timeout time.Duration) (string, error) {
	result := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second) // simulate slow function
		result <- "fetched data"
	}()

	select {
	case res := <-result:
		return res, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timeout after %v", timeout)
	}
}


func main(){
	d := make(chan int)
	b := make(chan int, 1)
	select {
	case v := <- d:
		fmt.Println("from a:", v)
	case b <- 42:
		fmt.Println("sent to b")
	case <- time.After(100 * time.Millisecond):
		fmt.Println("timeout")
	default:
		fmt.Println("nothing is ready right now")
	}

	// random choice
	count := make(map[int]int)
	a := make(chan int, 1000)
	c := make(chan int, 1000)
	for i := 0; i < 100; i++ {
		a <- 1
		c <- 0
	}
	for i := 0; i < 100; i++ {
		select {
		case v := <-a:
			count[v]++
		case v := <-c:
			count[v]++
		}
	}
	fmt.Println(count)

	// timeout and non-blocking send
	slow := make(chan int, 1)
	go func() {
		time.Sleep(200 * time.Millisecond)
		slow <- 1
	}()
	select {
	case r := <- slow:
		fmt.Println("got", r)
	case <- time.After(100 * time.Millisecond):
		fmt.Println("timeout adter 100ms")
	}
	// the gourotine behind slow is leaked, 
	// nobody is receiving from it, so it will block forever
	queue := make(chan int, 2)
	for i:=1; i <= 4; i++ {
		select {
		case queue <- i:
			fmt.Println("queued", i)
		default:
			fmt.Println("queue full, dropped", i)
		}
	}

	// merge two channels
	k := make(chan int)
	t := make(chan int)
	go func() {
		defer close(k)
		for i := 0; i < 10; i += 2 {
			k <- i
		}
	}()
	go func() {
		defer close(t)
		for i := 0; i < 10; i += 2 {
			t <- i + 1
		}
	}()
	merged := merge(k, t)
	for v := range merged {
		fmt.Print(v, " ")
	}
	fmt.Println()


	// producer
	words := []string{"hello", "world", "golang", "is", "awesome"}
	done := make(chan struct{})
	jobs := producer(words, done)
	fmt.Println(<-jobs)	// take only what we need
	close(done)			// signal the producer to stop
}