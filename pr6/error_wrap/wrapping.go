package main

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"strconv"
	"strings"
)

var ErrNotFound = errors.New("student not found")

func findStudent(name string) error{
	return ErrNotFound
}

func loadReport(name string) error{
	if err := findStudent(name); err != nil{
		return fmt.Errorf("load report for %q: %w", name, err)
	}
	return nil
}

func handler() error{
	if err := loadReport("Timur"); err != nil{
		return fmt.Errorf("GET /report: %w", err)
	}
	return nil
}

type GradeError struct{
	Student string
	Grade int
}
func (g *GradeError) Error() string{
	return fmt.Sprintf("invalid grade %d for %s", g.Grade, g.Student)
}

// own type can wrap too
type QueryError struct{
	Query string
	Err error
}
func (e *QueryError) Error() string { return "query " + e.Query + ": " + e.Err.Error()}
func (e *QueryError) Unwrap() error { return e.Err}		//without unwrap() we can't get into the Err inside the QueryError

var ErrTimeout = errors.New("timeout")


// tasks
// 1. Create GradeBook with methods Get(name) (int, error) (returns
// ErrNotFound ) and Set(name, grade) error (returns *GradeError for grades
// outside 0–100).
type GradeBook struct {
	grades map[string]int
}
func NewGradeBook() *GradeBook{
	return &GradeBook{
		grades: make(map[string]int),
	}
}
func (g *GradeBook) Get(name string) (int, error) {
	val, ok := g.grades[name]
	if !ok {
		return 0, ErrNotFound
	}
	return val, nil
}
func (g *GradeBook) Set(name string, grade int) error{
	if grade < 0 || grade > 100 {
		return &GradeError{name, grade}
	}
	g.grades[name] = grade
	return nil
}
// 2. Write a service function AddBonus(gb, name, bonus) error that calls Get
// and Set and wraps their errors with context.
func AddBonus(gb *GradeBook, name string, bonus int) error {
	val, err := gb.Get(name)
	if err != nil{
		return fmt.Errorf("get grade for %s: %w", name, err)
	}
	err = gb.Set(name, val+bonus)
	if err != nil{
		return fmt.Errorf("set grade for %s: %w", name, err)
	}
	return nil
}
// 3. Write report(err error) that prints OK , 404: ... (not found), 400: ...
// (invalid grade, also print the student and grade from the error) or 500:
// ... (anything else). Use errors.Is and errors.As.
func report(err error) {
	if err == nil {
        fmt.Println("OK")
        return
    }
	if errors.Is(err, ErrNotFound){
		fmt.Println("404:", err)
		return
	}
	var gradeErr *GradeError
	if errors.As(err, &gradeErr) {
		fmt.Printf("400: invalid grade %d for %s\n", gradeErr.Grade, gradeErr.Student)
		return
	}
	fmt.Println("500:", err)
}

// 4. parsing all the lines
// Write ParseAll(input string) ([]Record, error). A bad line must not stop the
// parsing: collect every error and return them together with errors.Join .
// Each error must say which line failed ( line 2: ... ) and keep the original error inside ( %w ).
// Missing comma → your own sentinel ErrBadFormat. Bad number → the error from strconv.Atoi .
type Record struct{
	name string
	grade int
}
var ErrBadFormat = errors.New("bad format")
func ParseAll(input string) ([]Record, error){
	var s []string = strings.Split(input, "\n")
	var r []Record
	var joined []error
	for i, v := range s {
		parts := strings.Split(v, ",")
		if len(parts) != 2{
			joined = append(joined, fmt.Errorf("line %d: %w", i+1, ErrBadFormat))
			continue
		}
		score, err := strconv.Atoi(strings.TrimSpace(parts[1]))
		if err != nil{
			joined = append(joined, fmt.Errorf("line %d: %w", i+1, err))
			continue
		}
		if score < 0 || score > 100 {
			joined = append(joined, fmt.Errorf("line %d: invalid grade", i+1))
			continue
		}
		r = append(r, Record{name: parts[0], grade: score})
	}
	return r, errors.Join(joined...)
}


func main(){
	err := handler()
	fmt.Println(err) 
	fmt.Println(err == ErrNotFound) //false
	fmt.Println(errors.Is(err, ErrNotFound)) //true

	// errors.Is(err,ErrNotFound) - Is this error or anything inside it ErrNotFound ?
	// errors.As(err, &target) - Is there an error of type *GradeError inside? Give it to me.
	// errors.Unwrap(err) - What is one layer deeper? 

	var wrapped error = fmt.Errorf("saving: %w", &GradeError{"F", 120})
	var ge *GradeError
	if errors.As(wrapped, &ge){
		fmt.Println("student:", ge.Student, "grade:", ge.Grade)
	}

	lost := fmt.Errorf("load report: %v", ErrNotFound)	//use %w instead of %v
	fmt.Println(lost)
	fmt.Println(errors.Is(lost, ErrNotFound))

	err = fmt.Errorf("load users: %w", &QueryError{Query: "Select * From users", Err: ErrTimeout})
	fmt.Println(errors.Is(err, ErrTimeout))

	// several error at once errrors.Join
	joined := errors.Join(ErrNotFound, &GradeError{"A", -6})
	fmt.Println(joined)
	fmt.Println(errors.Is(joined, ErrNotFound))	//true
	fmt.Println(errors.Join() == nil) 	//true: joining nothing gives nil

	// st lib
	_, err = os.Open("missing.txt")
	err = fmt.Errorf("config: %w", err)
	fmt.Println(errors.Is(err, fs.ErrNotExist))

	var pe *fs.PathError
	if errors.As(err, &pe){
		fmt.Println("operation:", pe.Op, "path:", pe.Path)
	}


	// tasks
	gb := NewGradeBook()
	report(gb.Set("A", 100))
	report(gb.Set("B", 50))
	report(gb.Set("C", -10))
	_, err = gb.Get("A")
	report(err)
	_, err = gb.Get("C")
	report(err)

	report(AddBonus(gb, "D", 10))
	report(AddBonus(gb, "A", 10))
	report(AddBonus(gb, "B", 30))

	// parsing
	input := `Aruzhan, 95
	Dias, 8O
	Timur
	Bota, 99999999999999999999
	Ali, 70`

	report, err := ParseAll(input)
	fmt.Println(report)
	fmt.Println(err)
	fmt.Println(errors.Is(err, ErrBadFormat))
	fmt.Println(errors.Is(err, strconv.ErrSyntax))
	fmt.Println(errors.Is(err, strconv.ErrRange))
}