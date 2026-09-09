package main

import (
	"fmt"
	"math"

	"github.com/fatih/color"
)

func change(a *int) {
	*a = 10000
}

func summ(a, b int) int {
	return a + b
}

func main() {
	fmt.Println("Hello, World!")
	var value int
	var value1 bool
	value2 := 10
	fmt.Println(value2)
	fmt.Println(value)
	fmt.Println(value1)

	first := 30
	second := 40
	res := summ(first, second)
	fmt.Println(res)

	third := 50
	change(&third)
	fmt.Println(third)

	forth := 4
	res2 := math.Sqrt(float64(forth))
	fmt.Println(res2)

	// Print with default helper functions
	color.Cyan("Prints text in cyan.")

	// A newline will be appended automatically
	color.Blue("Prints %s in blue.", "text")

	// These are using the default foreground colors
	color.Red("We have red")
	color.Magenta("And many others ..")
}
