package main

import (
	"fmt"
)

type Something struct{
	Some string
	Count int
}

func (s Something) add1 () {	//value receiver
	s.Count++
}

func (s *Something) add2 () {	//pointer receiver
	s.Count++
}

func main(){
	s1 := Something(Something{"Some", 10})
	s1.add1()
	fmt.Println(s1.Count)
	s1.add2()
	fmt.Println(s1.Count)
}