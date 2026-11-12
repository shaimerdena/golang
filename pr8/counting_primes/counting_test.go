package main

import (
	"testing"
	"fmt"
)

var sink int
const N = 200_000

func BenchmarkCountPar(b *testing.B) {
	for _, parts := range []int{2, 8, 1000, 200_000} {
		b.Run(fmt.Sprintf("parts=%d", parts), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				sink = CountPar(N, parts)
			}	
		})
	}
}

func BenchmarkCountSeq(b *testing.B) {
    for i := 0; i < b.N; i++ {
        sink = CountSeq(N)
    }
}

// Q: why does parts=2 on 2 cores not give a 2× speed-up, but parts=8
// comes closer? Why is parts=200000 the slowest of all — even slower than no
// goroutines?
// A: bc of the overhead of creating and managing goroutines. 
// With only 2 parts, the overhead may outweigh the benefits of parallelism, 
// especially if the workload is not large enough to justify it. 
// With 8 parts, the workload is better distributed across available cores, 
// leading to improved performance. However, with 200,000 parts, 
// the overhead of managing that many goroutines becomes significant, 
// resulting in slower performance than even a sequential approach.