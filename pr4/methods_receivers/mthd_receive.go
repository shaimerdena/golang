package main

import "fmt"

type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

func Area(r Rectangle) float64 {
	return r.Height * r.Width
}

type something int

func (s something) Sum() int {
	return int(s) + 1
}

//func (i int) Sum() int {    cannot define new methods on non-local type int
//	return i + 1
//}

func main() {
	r := Rectangle{Width: 10, Height: 5}
	fmt.Println(r.Area())
	fmt.Println(Rectangle.Area(r))
	fmt.Println(Area(r))

	s := something(12)
	fmt.Println(s.Sum())
}
