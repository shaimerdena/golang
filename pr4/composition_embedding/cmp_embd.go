package main

import (
	"fmt"
)

// composition
type Address struct{
	City string
}

type Person struct{
	Name string
	Age int
	Address Address	  //embedded
}

func (p Person) Greet() string {
	return "Hi, I am " + p.Name
}

func (e Employee) Greet() string{
	return e.Person.Greet() + " and I work at " + e.Company
}

// embedding
type Employee struct{
	Person		//embedded
	Company string
}

// name conflicts
type Engine struct{}
func (Engine) Start() string {  return "engine started"}

type Radio struct{}
func (Radio) Start() string { return "radio started" }

type Car struct {
	Engine
	Radio
}


type Animal struct{ Name string }
func (a Animal) Sound() string { return "..."}
func (a Animal) Describe() string { return a.Name + " says " + a.Sound() }

type Dog struct{ Animal }
func (d Dog) Sound() string { return "Woof!" }


// embedding a pointer
type Logger struct{ prefix string}
func (l *Logger) Log(msg string) { fmt.Println(l.prefix + msg)}

type Service struct{
	*Logger
	Name string
}

// tasks
// 1. Create Person with fields Name, Email and a method Contact() string
// that returns "Name <Email>".
type PersonT struct{
	Name string
	Email string
}
func (p PersonT) Contact() string {
	return p.Name + " <" + p.Email + ">" 
}

// 2. Create Student that embeds Person and adds GPA. Create Teacher that
// embeds Person and adds Department.
type Student struct{
	PersonT
	GPA float64
}
type Teacher struct{
	PersonT
	Department string
}

// 3. Give Teacher its own Contact() that also includes the department. Call
// Contact() on a student, on a teacher, and call the "hidden"
// Person.Contact() of the teacher.
func (t Teacher) Contact() string {
	return t.PersonT.Contact() + " <" + t.Department + ">"
}

// Participation Task: create a struct that embeds two types which both
// have a field ID. Show the compile error, then fix the code so that both IDs
// are printed.
type Phone struct{
	ID int
}
type Laptop struct{
	ID int
}
type Desk struct{
	Phone
	Laptop
}


func main(){
	p := Person{Name: "Yo", Age: 19, Address: Address{"Tokyo"}}
	fmt.Println(p.Address.City)

	e := Employee{
		Person: p,
		Company: "T",
	}
	fmt.Println(e.Person.Name)	//field promotion


	// shadowing
	//if the outer type declares a field/method 
	//with the same name, the outer wins 
	fmt.Println(e.Greet())			
	fmt.Println(e.Person.Greet())	//still reachable through the full path


	// name conflicts
	c := Car{}
	fmt.Println(c.Engine.Start()) 		// OK
	fmt.Println(c.Radio.Start()) 		// OK
	// fmt.Println(c.Start()) 				// compile error: ambiguous selector c.Start


	d := Dog{Animal{Name: "Rex"}}
	fmt.Println(d.Sound())		// Woof!
	fmt.Println(d.Animal.Sound())	//...
	fmt.Println(d.Describe())	//Rex says ..., bc Animal owns Describe method

	// embedding a pointer
	ok := Service{Logger: &Logger{prefix: "[svc] "}, Name: "users"}
	ok.Log("started")

	// bad := Service{Name: "orders"}
	// bad.Log("started")


	// tasks
	// 3.
	t := Teacher{Name: "D", Email: "d@gmail.com", Department: "site"}
	fmt.Println(t.Contact())
	sd := Student{Name: "J", Email: "j@outlook.com", GPA: 4.0}
	fmt.Println(sd.Contact())
	fmt.Println(t.PersonT.Contact())

	// 4.
	desk := Desk{Phone: Phone{1}, Laptop: Laptop{2}}
	// fmt.Println(desk.ID) 	//ambiguous selector desk.ID
	fmt.Println(desk.Phone.ID)
	fmt.Println(desk.Laptop.ID)
}