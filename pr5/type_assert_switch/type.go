package main

import (
	"errors"
	"fmt"
)

type Celsius float64
func (c Celsius) String() string {return fmt.Sprintf("%.1f°C", float64(c))}

// type switch
func describe(v any) string {
	// switch x := v.(type) {
	// case int, int64:
	// 	return x + 1		//won't compile bc type isn't asserted, it's any
	// }
	switch x := v.(type) {
	case nil:
		return "nil"
	case int, int64:
		return fmt.Sprintf("integer %v", x)
	case string:
		return fmt.Sprintf("string of %d bytes", len(x))
	case []int:
		return fmt.Sprintf("slice with %d ints", len(x))
	case error:
		return "error: " + x.Error()
	case fmt.Stringer:
		return "Stringer: " + x.String()
	case bool:
		return fmt.Sprintf("boolean: %v", x)
	case map[string]int:
		return fmt.Sprint(x)
	default:
		return fmt.Sprintf("unknown type %T: %v", x, x) 
	}
}

type Timeout struct { Sec int }
func (t Timeout) Error() string { return fmt.Sprintf("timeout after %ds", t.Sec)}
func (t Timeout) String() string { return "Timeout{" + fmt.Sprint(t.Sec) + "}"}
func kind(v any) string{
	switch v.(type){
	case fmt.Stringer:
		return "Stringer"
	case error:
		return "error"
	case Timeout:
		return "Timeout"
	}
	return "other"
}


// tasks
// Participation Task: rewrite describe without a type switch, using only v,
// ok := i.(T) assertions. Which version is easier to read?
func describe2(v any) string{
	if v == nil {
		return "nil"
	} 
	if value, ok := v.(int); ok{
		return fmt.Sprintf("integer: %v", value)
	} 
	if value, ok := v.(string); ok{
		return fmt.Sprintf("string: %v", value)
	} 
	if value, ok := v.(error); ok{
		return "error: " + value.Error()
	} 
	if value, ok := v.(fmt.Stringer); ok{
		return value.String()
	}
	return fmt.Sprintf("other type: %v", v)
}


func main(){
	// type assertation
	var i any = "Go"
	s := i.(string)
	fmt.Println(s)

	n, ok := i.(int)
	fmt.Println(n, ok)

	// f := i.(float64) panic: interface conversion: interface {} is string, not float64

	// asserting to another interface
	var t any = Celsius(26.6)
	if st, ok := t.(fmt.Stringer); ok{
		fmt.Println("Stringer: ", st.String())
	}

	// type switch
	fmt.Println(kind(Timeout{5}))
	var e error = Timeout{3}
	fmt.Println(e)	//timeout after 3s

	// tasks
	// 1. Copy describe from above and call it with: nil , 7 , int64(7) ,
	// "привет", []int{1, 2, 3} , errors.New("disk full") , Celsius(21.5) , 3.14 , struct{}{}.
	var x any
	fmt.Println(describe(x))
	var num1 any = 7
	fmt.Println(describe(num1))
	var num2 any = int64(7)
	fmt.Println(describe(num2))
	var str any = "привет"
	fmt.Println(describe(str))
	var list any = []int{1,2,3}
	fmt.Println(describe(list))
	var err any = errors.New("disk full")
	fmt.Println(describe(err))
	var cel any = Celsius(21.5)
	fmt.Println(describe(cel))
	var ft any = 3.14
	fmt.Println(describe(ft))
	var strct any = struct{}{}
	fmt.Println(describe(strct))
	var bln any = true
	fmt.Println(describe(bln))
	var mp any = map[string]int{"ali":5}
	fmt.Println(describe(mp))

	fmt.Println(describe2(num1))
}