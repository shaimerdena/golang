package main

import "fmt"

type Vertex struct{
	x, y int
}

func main(){
	map1 := make(map[string]int)
	p := &map1
	fmt.Println(*p)
	map1["alice"] = 19		//inserting and updating elements
	map1["some"] = 33
	fmt.Println(map1)

	vertex_map := map[string] Vertex{
		"Something": Vertex{12, 43},
		"Another something": {34, 54},
		"Also something": {x: 5, y: 9},
	}
	fmt.Println(vertex_map) 

	elem := vertex_map["Something"]
	fmt.Println(elem)		//retrieve an elem

	delete(vertex_map, "Something")		//deleting an elem
	fmt.Println(vertex_map)

	element, ok := vertex_map["Another something"]		//checking presence of an elem (ok returns true if an elem is present, and vice versa)
	fmt.Println(ok, element)
}