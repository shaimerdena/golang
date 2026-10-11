package main

import (
	"errors"
	"fmt"
	"io"
	"math"
	"os"
	"strings"
	"unicode/utf8"
)

type Shape interface{
	Area() float64
	Perimeter() float64
}

type Rect struct {W, H float64}
func (r Rect) Area() float64 { return r.H * r.W}
func (r Rect) Perimeter() float64 { return 2 * (r.H + r.W)}

type Circle struct {R float64}
func (c *Circle) Area() float64 { return math.Pi * c.R * c.R}
func (c *Circle) Perimeter() float64 { return 2 * math.Pi * c.R}

func printInfo(s Shape) {
	fmt.Printf("%-14T area=%6.2f perimeter=%6.2f\n", s, s.Area(), s.Perimeter())
}

func writeReport(w io.Writer, name string, score int){
	fmt.Fprintf(w, "%s: %d points\n", name, score)
}

// embedding interfaces
type Logger interface{ Log(msg string)}
type Saver interface{Save(data string) error}

type LogSaver interface{
	Logger 
	Saver
}

// tasks
// 1. Declare type Notifier interface { Notify(msg string) error }.
type Notifier interface{ Notify(msg string) error}
// 2. Implement it with three types: Email{Address} , SMS{Phone} and
// Telegram{Chat}. SMS must return an error if the message is longer than 20
// characters. Telegram must count how many messages it has sent (think
// about the receiver type!).
type Email struct{ Address string }
func (e Email) Notify(msg string) error{
	return nil
}

type SMS struct { Phone string }
func (s SMS) Notify(msg string) error{
	if(utf8.RuneCountInString(msg) > 20){
		return errors.New("the message is longer that 20 characters")
	}
	return nil
}

type Telegram struct { 
	Chat int
	Count int
 }
func (t *Telegram) Notify(msg string) error{
	t.Count++
	return nil
}
// Write Broadcast(ns []Notifier, msg string) []error that sends the
// message through every notifier and collects the errors.
func Broadcast(ns []Notifier, msg string) []error {
	error_list := make([]error, 0)
	for _, n := range ns{
		error_list = append(error_list, n.Notify(msg))
	}
	return error_list
}


func main(){
	shapes := []Shape{Rect{3,4}, &Circle{2}}
	for _, v := range shapes{
		printInfo(v)
	}

	// checking that a type implements an interface
	var _ Shape = Rect{}
	var _ Shape = (*Circle)(nil)

	var s Shape
	fmt.Printf("%v %T\n", s, s)	// <nil> <nil>
	s = Rect{1, 2}
	fmt.Printf("%v %T\n", s, s) // {1, 2} Rect
	s = &Circle{2}
	fmt.Printf("%v %T\n", s, s) // &{2} *(Circle)

	// any

	var x any
	x = 42
	x = "hello"
	x = []int{2,3}
	fmt.Println(x)

	// standard library interfaces
	// Interface		Method								Who implements it
	// fmt.Stringer		String() string						anything you want to print nicely
	// error			Error() string						every error value
	// io.Writer		Write(p Ilbyte) (n int, err error)	os. Stdout, files, *strings. Builder, HTTP responses
	// io. Reader		Read (p Ilbyte) (n int, err error)	os. Stdin, files, network connections

	writeReport(os.Stdout, "Aru", 90)	//print to the terminal
	var sb strings.Builder
	writeReport(&sb, "Dias", 80)	//write into a string
	fmt.Print(sb.String())


	// tasks
	// 4. Add compile-time checks ( var _ Notifier = ... ) for all three types
	var _ Notifier = Email{}
	var _ Notifier = SMS{}
	var _ Notifier = &Telegram{}


}