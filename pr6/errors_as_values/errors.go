package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

var grades map[string]int

// sentinel errors
var ErrNotFound = errors.New("student not found")

func GetGrade(name string) (int, error){
	g, ok := grades[name]
	if !ok {
		return 0, ErrNotFound
	}
	return g, nil
}

// custom error types
type GradeError struct{
	Student string
	Grade int
}
func (g *GradeError) Error() string{
	return fmt.Sprintf("invalid grade %d for %s", g.Grade, g.Student)
}
func SetGrade(name string, g int) error{
	if g < 0 || g > 100{
		return &GradeError{Student: name, Grade: g}
	}
	grades[name] = g
	return nil
}


func main(){
	// standard library
	n, err := strconv.Atoi("42a")
	fmt.Println(n, err)

	_, err = os.Open("missing.txt")
	fmt.Println(err)

	grades = map[string]int{
		"Mike": 50,
		"Ali": 90,
	}
	fmt.Println(GetGrade("Ciph"))

	fmt.Println(SetGrade("Mil", -9))
}