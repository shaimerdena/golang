package main

import (
	"fmt"
	"runtime"
	"runtime/debug"
	"bytes"
	"strings"
)

type Record struct {
	ID int
	Data [256]byte
}

var live []*Record //stays reachable: the live heap

func consume(r *Record) int {
	return r.ID
}
var sum int

func readFile(i int) []byte {
	data := make([]byte, 1<<20) //pretend to read 1 MiB of data from a file
	copy(data, fmt.Sprintf("HEADER-%04d", i))
	return data
}
func keepHeadersSlice(n int) [][]byte {
	var headers [][]byte
	for i:=0; i<n; i++ {
		data := readFile(i)
		headers = append(headers, data[:11]) // keep only the first 11 bytes (the header)
	}
	return headers
}
func keepHeadersCopy(n int) [][]byte {
	var headers [][]byte
	for i:=0; i<n; i++ {
		data := readFile(i)
		headers = append(headers, bytes.Clone(data[:11])) // keep only the first 11 bytes (the header)
	}
	return headers
}

func readFileString(i int) string {
	return strings.Repeat("A", 1<<20)
}
func keepStringSlice(n int) []string {
	var headers []string
	for i:=0; i<n; i++ {
		data := readFileString(i)
		headers = append(headers, data[:5]) // keep the string
	}
	return headers
}
func keepStringCopy(n int) []string {
	var headers []string
	for i:=0; i<n; i++ {
		data := readFileString(i)
		headers = append(headers, strings.Clone(data[:5])) // keep the string
	}
	return headers
}

func main(){
	// mark and sweep
	// - marks reachable objects; sweeps unreachable objects
	// concurrent
	// - most of the time, the program runs concurrently with the garbage collector
	// generational
	// - young generation: short-lived objects, collected frequently
	// - old generation: long-lived objects, collected less frequently
	// non-moving
	// - objects are not moved in memory, which avoids the need for updating references
	// write barrier
	// - a mechanism to keep track of changes to objects during garbage collection

	// the GC starts a new cycle when the heap size exceeds the heap goal
	// heap goal = live heap * (1 + GOGC/100)

	for i:=0; i<80000; i++ {
		live = append(live, &Record{ID: i})
	}
	for i:=0; i < 10_000_000; i++{
		r := &Record{ID: i}	//garbage after one use
		sum += consume(r)
	}

	// Q: With GOGC=off there is no GC at all — so why is it the slowest run?
	// A: bc the huge amount of garbage results in a lot of memory being allocated,
	// which can slow down the program and degrade perfomance.

	// Q: And why does GOMEMLIMIT=100MiB together with GOGC=off work so
	// well here?
	// A: bc GOGC=off disables the GC, but GOMEMLIMIT=100MiB sets a soft limit on memory usage. 
	// This prevents the program from consuming too much memory, which helps maintain performance.
	// TIP: In containers (Docker, Kubernetes) set GOMEMLIMIT a bit below the container's memory limit (e.g. 90 %).

	debug.SetMemoryLimit(500 >> 20)	// set memory limit to 500 MiB
	// altenative - GOMEMLIMIT=500MiB go run gc.go
	debug.SetGCPercent(100) //set GOGC to 100% (default value)
	// alternative - GOGC=100 go run gc.go

	// GOGC=off disables the garbage collector, which can lead to increased memory usage and potential performance degradation.
	runtime.GC() // force a garbage collection cycle


	keepHeadersSlice(1000)
	keepHeadersCopy(1000)

	keepStringSlice(1000)
	data := keepStringCopy(1000)


	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Println("Number of allocations:", m.Mallocs)
	fmt.Println("Number of frees:", m.Frees)
	fmt.Println("Heap in use:", m.HeapAlloc>>10, "KB")
	fmt.Println("Number of GC cycles:", m.NumGC)
	fmt.Println("Sum of consumed values:", sum)
	runtime.GC()
	runtime.ReadMemStats(&m)
	runtime.KeepAlive(data) // keep data reachable until this point

	// Q: A slice data[:11] has len 11 — why does it keep the whole 1 MB
	// array alive? (Week 2: what's inside a slice header?)
	// A: A slice header contains a pointer to the underlying array, the length of the slice, and the capacity of the slice. 
	// When you create a slice like data[:11], 
	// it still points to the original 1 MB array in memory. 
	// As long as that slice is reachable, the entire underlying array remains reachable 
	// and cannot be garbage collected, even if you only need the first 11 bytes.
}