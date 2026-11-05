package main

import (
	"fmt"
	"runtime"
)

type Record struct {
	ID int
	Data [256]byte
}

var live []*Record

func consume(r *Record) int {
	return r.ID
}

var sum int

func main(){
	for i:=0; i < 80_000; i++ {
		live = append(live, &Record{ID:i})
	}
	for i:=0; i < 10_000_00; i++ {
		r := &Record{ID:i}
		sum += consume(r)
	}
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	runtime.GC()
	fmt.Println("Number of allocations:", m.Mallocs)
	fmt.Println("Number of frees:", m.Frees)
	fmt.Println("Heap in use:", m.HeapAlloc>>10, "KB")
	fmt.Println("Number of GC cycles:", m.NumGC)
	fmt.Println("Sum of consumed values:", sum)

	// | GOGC | Allocations | Frees | Heap in use | GC cycles | Sum of consumed values |
	// | 50   | 80211       | 58    | 24468 KB     | 7         | 499999500000           |
	// | 100  | 80217       | 57    | 24849 KB     | 3         | 499999500000           |
	// | 200  | 80209       | 59    | 23995 KB     | 2         | 499999500000           |
	// | 400  | 80209       | 57    | 24843 KB     | 1         | 499999500000           |
	// | off  | 80177       | 6     | 26106 KB     | 0         | 499999500000           |

	// gc 9 @0.031s 4%: 0.018+0.26+0.002 ms clock, 0.14+0.010/0.36/0.58+0.023 ms cpu, 3->3->1 MB, 4 MB goal, 0 MB stacks, 0 MB globals, 8 P
	// 9 gc cycles in 0.031s, 4% of total time, 0.018+0.26+0.002 ms clock, 0.14+0.010/0.36/0.58+0.023 ms cpu, 3->3->1 MB, 4 MB goal, 0 MB stacks, 0 MB globals, 8 P

	// GOGC=off GOMEMLIMIT=100MiB go run gc.go
	// Number of GC cycles: 0
	// memory is less than 100MiB, so no GC cycles are triggered. The program runs faster and uses less memory.
}