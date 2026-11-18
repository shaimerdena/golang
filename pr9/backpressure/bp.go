package main

import (
	"time"
	"errors"
)

var ErrQueueFull = errors.New("queue is full")

func Submit(queue chan <- string, job string, timeout time.Duration) error {
	t := time.NewTimer(timeout)
	defer t.Stop()
	
	select {
	case queue <- job:
		return nil
	case <-t.C:
		return ErrQueueFull
	}
}

func main() {
	// consumer is slower than producer
	// the pressure "pushes back" along the pipeline
	// and slows down the producer

	// strategies to handle backpressure:
	// 1. huge buffer: the consumer can catch up later
	// 2. drop: the producer drops data when the consumer is slow
	// 3. block: the producer blocks when the consumer is slow
	// 4. wait but not forever: the producer waits for a while, then drops data or returns an error

	// Q: The "huge buﬀer" handled everything and the producer was fast —
	// so why is it the worst choice for a web server? Think of a job as an
	// HTTP request that a user is waiting for. What happens if the
	// producer never stops?
	// A: If the producer never stops and uses a huge buffer, 
	// it can lead to unbounded memory usage, which can cause the web server to run out of memory and crash. 
	// Additionally, users may experience long wait times for their requests to be processed, 
	// leading to poor user experience.
}