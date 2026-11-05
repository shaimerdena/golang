package main

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
)

type User struct {
	Name string
}
func MyFunctionPointer() *User {
    u := &User{Name: "Patrick Jane"}
    return u
}
func MyFunctionValue() User {
    u := User{Name: "Patrick Jane"}
    return u
}

func JoinIDsSlow(ids []int) string {
	s := ""
	for i, id := range ids {
		if i > 0 {
			s += ","
		}
		s += fmt.Sprint(id)	// This line creates a new string for each id, leading to multiple allocations.
	}
	return s
}
func JoinIDsFast(ids []int) string {
	var b strings.Builder
	b.Grow(len(ids) * 8) // Preallocate capacity to reduce allocations
	var tmp [20]byte // Temporary buffer for integer to string conversion
	for i, id := range ids {
		if i > 0 {
			b.WriteByte(',')
		}
		b.Write(strconv.AppendInt(tmp[:0], int64(id), 10))
	}
	return b.String()
}
// JoinIDsFast still makes one allocation. Which one, and why can't it
// be avoided?
// A: bc it returns a string, which is a slice of bytes.
// The allocation cannot be avoided because the size of the resulting string is not known at compile time, 
// so it must be allocated dynamically at runtime. 

func CollectNoCap(n int) []int {
	var numbers []int
	for i := 0; i < n; i++ {
		numbers = append(numbers, i)
	}
	return numbers
}
func CollectWithCap(n int) []int {
	numbers := make([]int, 0, n)
	for i := 0; i < n; i++ {
		numbers = append(numbers, i)
	}
	return numbers
}
// Why does CollectNoCap(1000) make exactly 12? (Hint: Week 2, how cap grows.)
// A: Because the slice grows dynamically as elements are appended, and the capacity doubles each time it runs out of space. 
// Starting from an initial capacity of 0, the capacities will be: 1, 2, 4, 8, 16, 32, 64, 128, 256, 512, 1024. 
// This results in a total of 12 allocations (1 for each capacity increase).


func main(){
	u := MyFunctionPointer()
	println(u.Name)

	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Println("Number of allocations:", m.Mallocs)
	fmt.Println("Number of frees:", m.Frees)
	fmt.Println("Heap in use:", m.HeapAlloc>>10, "KB")
	fmt.Println("Number of GC cycles:", m.NumGC)

	// Typical sources of allocations                 // Possible optimizations

	// Returning pointers to small structs             -> Consider returning values instead
	// Repeated string concatenation in loops          -> Use strings.Builder or bytes.Buffer
	// append() without knowing/preallocating capacity -> Preallocate slice capacity when possible
	// fmt.Sprint()/fmt.Sprintf() when writing to Writer -> Use fmt.Fprint()/fmt.Fprintf()
	// Unnecessary interface conversions (any)         -> Avoid unnecessary boxing when possible
	// Creating a new buffer for every operation       -> Reuse buffers; consider sync.Pool
	// Repeated []byte <-> string conversions          -> Avoid unnecessary conversions


}