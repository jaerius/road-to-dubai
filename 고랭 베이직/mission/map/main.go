package main

import (
	"fmt"
)

func main() {
	var m1 map[string]int
	fmt.Println(m1)

	m2 := map[string]int{"a": 1, "b": 2}
	fmt.Println(m2)

	m3 := make(map[string]int)
	fmt.Println(m3)

	m3["a"] = 1
	fmt.Println(m3)
	m3["a"] = 10
	fmt.Println(m3)

	val := m3["a"]
	fmt.Println(val)

	val2 := m3["b"]
	fmt.Println(val2)

	delete(m3, "a")
	fmt.Println(m3)

}
