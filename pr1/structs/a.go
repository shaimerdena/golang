package main

import "fmt"

type Vertex struct{
	X, Y int
}

var (
	v1 = Vertex{1,2}
	v2 = Vertex{X: 10}
	v3 = Vertex{}
	v4 = &Vertex{15, 14}
)

func main(){
	v1.X = 7
	fmt.Println(v1)
	p := &v2
	p.Y = 10
	fmt.Println(*p)
	fmt.Println(v3)
	v4.Y = 9
	fmt.Println(*v4)
}