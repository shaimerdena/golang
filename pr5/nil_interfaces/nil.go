package main

import (
	"errors"
	"fmt"
)

type Shape interface{ Area() float64}

type Square struct{ Side float64}
func (s *Square) Area() float64{
	if s == nil {
		return 0
	}
	return s.Side * s.Side
}

//tasks
type ValidationError struct { Field string }
func (e *ValidationError) Error() string { return "invalid field: " + e.Field}
func validate(age int) error{
	var err *ValidationError
	if age < 0{
		err = &ValidationError{Field: "age"}
		return err
	}
	return nil
}

func main(){
	var s Shape
	fmt.Println(s == nil) //true

	var sq *Square
	s = sq
	fmt.Println(s == nil) //false
	fmt.Println(s.Area()) //0

	// s = nil
	// fmt.Println(s.Area()) //doesn't work
	// Why can Go call Area() when the interface
	// holds a nil pointer, but not when the interface itself is nil ?
	// A: bc it has no dynamic type which could run the function

	err := validate(-1)
	var ve *ValidationError
	if errors.As(err, &ve){
		fmt.Println("invalid field:", ve.Field)
	}
}