package main

import (
	"fmt"
)

type Student struct{
		Name string
		ID int
		GPA float64
		Courses []string
	}

type Team struct {
	Name string
	Scores [3]int
	Members []string
}

func addCourse(s Student, course string){
		s.Courses = append(s.Courses, course)
	}

func addCourseChanged(s *Student, course string){
		s.Courses = append(s.Courses, course)
	}

func main(){
	// student example
	
	var s1 Student	//creating struct with default values
	fmt.Println(s1)
	s2 := Student{"Dias", 123, 4.0, []string{"OOP", "ADS"}}		//positional arguments
	fmt.Println(s2)
	s3 := Student{Name: "Alina", ID: 124}	//keyword arguments
	fmt.Println(s3)
	s4 := &Student{Name: "c"} 
	s4.ID = 125
	s5 := new(Student)	//pointer to a zero Student
	fmt.Println(s5)

	fmt.Println("-------------")

	fmt.Printf("%v\n", s1)		//default format
	fmt.Printf("%+v\n", s1)		//default but with field names
	fmt.Printf("%#v\n", s1)		//go-syntax format
	fmt.Println(s1.Courses == nil)

	// team example
	// copies all of its fields

	// Why did one field change in both structs, while the others didn't?
	// bc copying a slice field copies only the slice header (copies pointer)
	t1 := Team{Name: "Gophers", Scores: [3]int{1, 2, 3}, Members: []string{"Ali", "Bota"}}
	t2 := t1 // copy
	t2.Name = "Rustaceans"
	t2.Scores[0] = 100
	t2.Members[0] = "Nurlan"
	fmt.Println(t1.Name, t1.Scores, t1.Members)
	fmt.Println(t2.Name, t2.Scores, t2.Members)

	// comparing structs

	type Point struct{ X, Y int }
	a, b := Point{1, 2}, Point{1, 2}
	c, d := &Point{1, 2}, &Point{1, 2}
	fmt.Println(a == b) // true
	fmt.Println(c == d) // false - bc it's comparing addresses
	fmt.Println(*c == *d) // true
	
	// fmt.Println(t1 == t2) // cannot be compared bc of having not comparable field ([]string)

	// anon / empty

	fmt.Println("---------")
	point := struct{ X, Y int }{1, 2} // one-off struct without a type name
	fmt.Println(point)
	seen := map[string]struct{}{} // a "set"
	seen["go"] = struct{}{}
	_, ok := seen["go"] // true
	fmt.Println(ok)


	
	// tasks
	// 1. Declare the Student struct above and create students in 4 diﬀerent
	// ways (zero value, named fields, positional, pointer). Print each one with %+v.

	var sd1 Student
	sd2 := Student{Name: "cipher", ID: 1, GPA: 4.0}
	sd3 := Student{"cipher", 2, 4.0, []string{"OOP"}}
	sd4 := &Student{Name: "cipher", ID: 3, GPA: 4.0}
	arr := []Student{
		sd1, sd2, sd3, (*sd4),
	}

	for i, sd := range arr{
		fmt.Println(i, sd)
	}

	// 2. Write func addCourse(s Student, course string) that appends a course to
	// s.Courses. Call it and print the student. Did the course appear? Why?

	addCourse(sd2, "ADS")
	fmt.Println(sd2)

	// 3. Change the function to take *Student and call it again. What changed?
	addCourseChanged(&sd2, "ADS")
	fmt.Println(sd2)
}	