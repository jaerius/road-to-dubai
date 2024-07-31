package main

import (
	"fmt"
)

type Speaker interface {
	Speak() string
}

type Person struct {
	name string
}

func (p Person) Speak() string {
	return "hi my name is" + p.name
}

type Dog struct {
	name string
}

func (d Dog) Speak() string {
	return "walwal my name is" + d.name
}

func Greet(s Speaker) {
	fmt.Println(s.Speak())
}

func main() {
	p := Person{name: "jaewon"}
	d := Dog{name: "gae"}

	Greet(p)
	Greet(d)
}
