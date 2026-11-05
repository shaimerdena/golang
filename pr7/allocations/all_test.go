package main

import (
	"fmt"
	"testing"
)

var userSink *User
var valueSink User

func BenchmarkPointer(b *testing.B) {
    for i := 0; i < b.N; i++ {
        userSink = MyFunctionPointer()
    }
}

func BenchmarkValue(b *testing.B) {
    for i := 0; i < b.N; i++ {
        valueSink = MyFunctionValue()
    }
}

func TestAllocations(t *testing.T) {
	allocs := testing.AllocsPerRun(1000, func() { userSink = MyFunctionPointer() })
	fmt.Println("Allocations: ", allocs)
}

func TestAllocationsValue(t *testing.T) {
	allocs := testing.AllocsPerRun(1000, func() { valueSink = MyFunctionValue() })
	fmt.Println("Allocations: ", allocs)
}

// tasks
// 1. Copy JoinIDsSlow and write a test with testing.AllocsPerRun that prints
// its allocations for 10 IDs.
func TestJoinIDsSlow(t *testing.T) {
	allocs := testing.AllocsPerRun(1000, func() { _ = JoinIDsSlow([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}) })
	fmt.Println("Allocations: ", allocs)
}

// 2. Write JoinIDsFast and a test that fails if it makes more than 1
// allocation.
func TestJoinIDsFast(t *testing.T) {
	allocs := testing.AllocsPerRun(1000, func() { _ = JoinIDsFast([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}) })
	
	if allocs > 1 {
		t.Errorf("Expected 1 allocation, got %f", allocs)
	}
}

// Write BenchmarkJoinIDsSlow and BenchmarkJoinIDsFast and compare ns/op ,
// B/op , allocs/op.
func BenchmarkJoinIDsSlow(b *testing.B) {
	ids := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i := 0; i < b.N; i++ {
		_ = JoinIDsSlow(ids)
	}
}
func BenchmarkJoinIDsFast(b *testing.B) {
	ids := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i := 0; i < b.N; i++ {
		_ = JoinIDsFast(ids)
	}
}

func BenchmarkCollectNoCap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = CollectNoCap(1000)
	}
}

func BenchmarkCollectWithCap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = CollectWithCap(1000)
	}
}