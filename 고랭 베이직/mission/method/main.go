package main

import (
	"fmt"
)

type Rectangle struct {
	width, height float64
}

func (r Rectangle) area() float64 {
	return r.width * r.height
}

func (r *Rectangle) scale(factor float64) {
	r.width *= factor
	r.height *= factor
}

func main() {
	rect := Rectangle{width: 3, height: 4}
	fmt.Println("Area:", rect.area())

	rect.scale(2)
	fmt.Println("new dimensions:", rect.width, rect.height)
	fmt.Println("new area:", rect.area())

}
