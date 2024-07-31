package main

import (
	"fmt"
)

type Person struct {
	name string
	age  int
}

func main() {
	p := Person{name: "Alice", age: 30}

	fmt.Println(p.name)
	fmt.Println(p.age)

	p.age = 31
	fmt.Println(p.age)
}
