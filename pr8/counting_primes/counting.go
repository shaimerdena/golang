package main

import (
	"sync"
	"sync/atomic"
	"fmt"
)

func isPrime(n int) bool {
	for i := 2; i * i <= n; i++ {
		if n % i == 0{
			return false
		}
	}
	return true
}

func CountSeq(n int) int {
	count := 0
	for i := 0; i < n; i++ {
		if isPrime(i) {
			count++
		}
	}
	return count
}

func CountPar(n, parts int) int {
	var wg sync.WaitGroup
    var count atomic.Int64
    part := n / parts

    for i := 0; i < parts; i++ {
        start := i * part
        end := start + part

        if i == parts-1 {
            end = n
        }

        wg.Add(1)

        go func(start, end int) {
            defer wg.Done()

            for j := start; j < end; j++ {
                if isPrime(j) {
                    count.Add(1)
                }
            }
        }(start, end)
    }

    wg.Wait()

    return int(count.Load())
}

func main() {
    n := 200000

    seq := CountSeq(n)
    par := CountPar(n, 4)

    fmt.Println("Sequential:", seq)
    fmt.Println("Parallel:", par)
}