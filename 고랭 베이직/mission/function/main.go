package main

import (
	"fmt"
)

func greet() {
	fmt.Println("Hello, World!")
}

func add(a int, b int) int {
	return a + b
}

func main() {
	greet()
	result := add(3, 4)
	fmt.Println("Sum:", result)
}
