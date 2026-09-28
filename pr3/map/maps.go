package main

import (
	"fmt"
)

func check_existence(m map[string]int, key string) {
	_, ok := m[key]
	if(ok == true){
		fmt.Println(key, "is in the map")
	} else{
		fmt.Println(key, "is missing")
	}
} 

func main(){
	// How to test whether a key is present with a two-value assignment?
	m := map[string]int {
		"ayau": 10,
		"cipher": 40,
		"someone": 15,
	}
	delete(m, "someone")
	val, ok := m["someone"]
	fmt.Println(val, ok)
	
	// map init
	p := make(map[string]int, 100)
	fmt.Println(len(p))

	// student exam scores
	scores := map[string]int{
		"Alice": 90,
		"Bob": 0,
	}
	check_existence(scores, "Bob")
	check_existence(scores, "Charlie")

	// inventory management
	inventory := map[string]int{
		"apples": 10,
		"bananas": 5,
	}
	fmt.Println("apples - ", inventory["apples"])
	inventory["bananas"] += 7
	inventory["oranges"] = 8
	delete(inventory, "apples")
	fmt.Println(inventory)
}