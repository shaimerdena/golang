package main

import (
	"bytes"
	"fmt"
	"strings"
	"unicode/utf8"
)

func concat(values []string) string {
		var sb strings.Builder
		for _, v := range values {
			sb.WriteString(v)
		}
		return sb.String()
	}

// old version
func processPayload(data []byte) string {
	// Unnecessary type conversions
	trimmed := string(bytes.TrimSpace(data))
	clean := strings.ReplaceAll(trimmed, "\n","")
	return clean
}

// refactored version
func processPayload2(data []byte) []byte{
	trimmed := bytes.TrimSpace(data)
	clean := bytes.ReplaceAll(trimmed, []byte("\n"), []byte(""))
	return clean
}

func main() {
	s := "Hello ворлд"
	fmt.Println(len(s), s) //prints a size in bytes
	fmt.Println(s[7])
	for i, r := range s{
		fmt.Printf("position %d: %c\n", i, r)
	}

	s = "世界"
	fmt.Println(len(s))
	fmt.Println(utf8.RuneCountInString(s))

	runes := []rune(s)
	fmt.Printf("%c\n", runes[1])

	// Trim functions
	right := "oxo123oxo"
	left := "xo"
	fmt.Println(strings.TrimLeft(right,left))
	fmt.Println(strings.TrimRight(right,left))
	fmt.Println(strings.TrimSuffix(right,left))
	fmt.Println(strings.TrimPrefix(right,left))

	// concatenation
	list_of_words := []string{"one", "two", "three"}
	concat_words := concat(list_of_words)
	fmt.Println(concat_words)

	// grow
	var b strings.Builder
	b.Grow(1024)
	for i:=0; i<50; i++{
		b.WriteString("Something")
	}
	fmt.Println("Len: ", len(b.String()))

	str := " Hello "
	str_bytes := []byte(str)
	str_bytes = bytes.TrimSpace(str_bytes)
	fmt.Printf("%c\n", str_bytes)


	// Explain what happens under the hood in terms of memory
	// allocation when converting between []byte and string.
	// Ans: an array of bytes translate into slice of bytes

	// Refactor the processPayload function to work entirely with
	// byte slices using the standard bytes package, eliminating
	// unnecessary type conversions.

	g := " skfld fldks\n " 
	k := processPayload([]byte(g))
	fmt.Print(k)
	fmt.Println("something")
}
