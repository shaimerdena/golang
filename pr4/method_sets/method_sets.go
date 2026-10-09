package main

import (
	"fmt"
	"math"
)

type Speaker interface {
	Speak() string
}
type Cat struct{ Name string }
func (c Cat) Speak() string { return c.Name + ": meow" }

type Dog struct{ Name string }
func (d *Dog) Speak() string { return d.Name + " : woof" }



type Speaker_ interface{ Speak() string }

type Base struct{}

func (Base) Hello() string  { return "hello" }
func (*Base) Speak() string { return "speak" }

type ByValue struct{ Base }    // embeds Base
type ByPointer struct{ *Base } // embeds *Base



type Book struct {
	Title  string
	Author string
}
func (b *Book) String() string { return b.Title + " by " + b.Author }



// tasks
// 1. Declare type Shape interface { Area() float64; Perimeter() float64 }.
// 2. Implement Rectangle with value receivers, plus a method Scale(k
// float64) that really changes the rectangle.
// 3. Implement Circle with pointer receivers.

type Shape interface{
	Area() float64
	Perimeter() float64
}

type Rectangle struct{
	Width float64
	Length float64
}
func (r Rectangle) Area() float64 {
	return r.Length * r.Width
}
func (r Rectangle) Perimeter() float64 {
	return 2*(r.Length + r.Width)
}
func (r *Rectangle) Scale(k float64){
	r.Width *= k
	r.Length *= k
}

type Circle struct{
	Radius float64
}
func (c *Circle) Area() float64{
	return math.Pi * c.Radius * c.Radius
}
func (c *Circle) Perimeter() float64{
	return 2 * c.Radius * math.Pi
}



type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.


func main() {
	var s Speaker

	s = Cat{"Murka"}
	s = &Cat{"Felix"}
	s = &Dog{"Tarzan"}
	// s = Dog{"Bim"}	| doesn't compile bc method speak has pointer receiver
	fmt.Println(s.Speak())

	// We just saw that d.Speak() works on a Dog variable (Go adds &
	// for us). So why can't a Dog value be put into a Speaker ?

	// ans: d.Speak() is addressable. Go calls (&d).Speak() for us.
	// but when Dog value is stored in an interface, the interface stores copy of that value (not addresable)
	// so it can't do (&Dog) for us. only &Dog satisfies Speaker.

	var _ Speaker = (*Dog)(nil) //fails to compile if *Dog stops implementing Speaker
	var _ Speaker = Cat{}

	// Struct S embeds… method set of S 		method set of *S
	// T 				methods of T 			methods of T and *T
	// *T 				methods of T and *T 	methods of T and *T

	var _ Speaker_ = &ByValue{} // compiles
	// var _ Speaker = ByValue{} 	// no | bc
	var _ Speaker_ = ByPointer{Base: &Base{}}  // compiles
	var _ Speaker_ = &ByPointer{Base: &Base{}} // compiles

	b := Book{"The Go Programming Language", "Donovan & Kernighan"}
	fmt.Println(b)
	fmt.Println(&b)


	// tasks
	var sr Shape
	sr = Rectangle{4, 8}
	fmt.Println(sr.Area(), sr.Perimeter())
	r := Rectangle{10, 20}
	r.Scale(0.5)
	fmt.Println(r)

	shapes := []Shape{
		Rectangle{4, 5},
		&Circle{5},
	}
	for _, v := range shapes{
		fmt.Println(v.Area(), v.Perimeter())
	}
	// shapes = append(shapes, Circle{7})
	// compiler error:cannot use Circle{…} (value of struct type Circle) as 
	// Shape value in argument to append: 
	// Circle does not implement Shape (method Area has pointer receiver)


	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}

	rct := Rectangle{3, 4}
	f := rct.Area // a "method value"
	rct.Width = 10
	fmt.Println(f(), rct.Area()) // 12, 40
}
