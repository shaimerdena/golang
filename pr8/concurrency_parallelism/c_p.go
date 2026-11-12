package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func countPrimes(start, end int) int {
	count := 0
	for i := start; i < end; i++ {
		if isPrime(i) {
			count++
		}
	}
	return count
}

// cpu-bound: each "task" is a prime number check
func cpuWork(workers int) (int, time.Duration){
	const iterations = 6_000_000
	start := time.Now()
	results := make([]int, workers)
	var wg sync.WaitGroup
	chunk := iterations / workers
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			results[i] = countPrimes(i*chunk, (i+1)*chunk)
		}(i)
	}
	wg.Wait()
	total := 0
	for _, v := range results {
		total += v
	}
	return total, time.Since(start)
}
// i/o-bound: each "request" just waits 100 ms
func ioWork(workers int) time.Duration {
	start := time.Now()
	var wg sync.WaitGroup
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			time.Sleep(100 * time.Millisecond)	// Simulate I/O work
		}()
	}
	wg.Wait()
	return time.Since(start)
}

func main(){
	cpuWorkers := 10
	cpuCount, cpuDuration := cpuWork(cpuWorkers)
	println("CPU-bound work:")
	println("Number of primes found:", cpuCount)
	println("Time taken:", cpuDuration.String())

	ioWorkers := 10
	ioDuration := ioWork(ioWorkers)
	println("\nI/O-bound work:")
	println("Time taken for", ioWorkers, "workers:", ioDuration.String())

	// | Gouroutines | CPU-bound time | I/O-bound time |
	// |-------------|----------------|----------------|
	// | 1           | 573.814125ms   | 101.054917ms   |
	// | 2           | 340.636959ms   | 101.106875ms   |
	// | 4           | 209.988334ms   | 101.110167ms   |
	// | 8           | 161.814292ms   | 101.127209ms   |
	// | 16          | 164.490458ms   | 101.073291ms   |

	fmt.Println(runtime.NumCPU(), "CPU cores available")

	// GOMAXPROCS can be set to control the number of OS threads used for CPU-bound work
	// GOMAXPROCS=2 go run c_p.go

	// questions:
	// 1. With GOMAXPROCS=1 , why doesn't the CPU work get faster with
	// more goroutines, but 1000 "requests" still finish in 100 ms?
	// A: bc every goroutine just waits for the I/O operation to complete.

	// 2. With 2 cores, 4 goroutines are not faster than 2. Why? And why
	// is 2 goroutines not exactly 2× faster than 1? (Hint: are bigger
	// numbers harder to check?)	
	// A: bc the max number of threads is 2, so 4 goroutines will be scheduled on 2 threads. 
	// Also, bigger numbers take longer to check for primality, so the workload is not evenly distributed.

	runtime.GOMAXPROCS(1)
	var wg sync.WaitGroup
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Print(i, " ")
		}()
	}
	wg.Wait()
	// Q: What does it print, and is it always the same? The Go specification promises
	// nothing about the order of goroutines — so why is it a bug to rely on this
	// order, even if it never changes on your laptop?
	// A: prints numbers in random order, bc goroutines are scheduled independently. 
	// Relying on this order is a bug because it can lead to non-deterministic behavior, 
	// making the program unpredictable and harder to debug,
	// especially when run on different machines or under different conditions.
}