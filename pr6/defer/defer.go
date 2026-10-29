package main

import (
	"fmt"
	"os"
	"time"
)

func double() (result int){
	defer func() {result*=2}()
	return 21
}

func notDouble() int{
	result := 21
	defer func() {result*=2}()
	return result
}

func save(name string) (err error){
	f, err := os.Create(name)
	if err != nil{
		return err
	}
	defer func() {
		if cerr := f.Close(); cerr != nil && err == nil{
			err = cerr
		}
	}()
	_, err = f.WriteString("Aruzhan,95\n")
	return err
}

func trace(name string) func(){
	fmt.Println("enter", name)
	return func() { fmt.Println("exit", name) }
}
func work() {
	defer trace("work")()
	fmt.Println("working...")
}
// enter work
// working...
// exit work


// tasks
// 2. Write a fake Open(name) (*Resource, error) and (*Resource).Close() that
// print open / close. Write processAllBad(names) with defer inside the loop
// and processAllGood(names) that fixes it with a helper function. Show the
// diﬀerence in the order of open / close .
type Resource struct {
	Name string
}
func Open(name string) (*Resource, error) {
	fmt.Println("open", name)
	return &Resource{}, nil
}
func (r *Resource) Close() error {
	fmt.Println("close", r.Name)
	return nil
}
func processAllBad(names []Resource) {
	for _, v := range names {
		Open(v.Name)
	}
	for _, v := range names{
		defer v.Close()
	}
}
func helper(r *Resource){
	defer r.Close()
	Open(r.Name)
}
func processAllGood(names []Resource) {
	for _, v := range names {
		helper(&v)
	}
}
// 3. Use the trace trick to print how long a function 
// took: defer timer("sort")() , with time.Since.
func timer(name string) func() {
	start := time.Now()
	return func() {
		fmt.Println(name, time.Since(start))
	}
}
func sort() {
	defer timer("sort")()
	time.Sleep(1 * time.Second)
	fmt.Println("sorting...")
}


func main(){
	// defer
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
	fmt.Println("main")
	// main 3 2 1

	x := 1
	defer fmt.Println("deferred x=", x)
	defer func() { fmt.Println("closure x =", x) }()
	x = 2
	fmt.Println("current x =", x)
	// curr = 2
	// closure = 2
	// defer = 1
	
	fmt.Println(double(), notDouble()) // 42 21
	
	fmt.Println(save("sample.txt"))

	work()

	fmt.Println()
	processAllBad([]Resource{{Name: "file1"}, {Name: "file2"}})
	processAllGood([]Resource{{Name: "file1"}, {Name: "file2"}})
	fmt.Println()

	sort()
}