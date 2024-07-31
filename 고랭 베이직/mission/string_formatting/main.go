package main

import (
	"fmt"
)

func main() {
	intVar := 123
	floatVar := 123.456
	strVar := "Hello, World!"
	boolVar := true
	pointerVar := &intVar

	fmt.Printf("Default format%v\n", intVar)
	fmt.Printf("Default format%v\n", floatVar)
	fmt.Printf("Default format%v\n", strVar)
	fmt.Printf("Default format%v\n", boolVar)
	fmt.Printf("Default format%v\n", pointerVar)

	fmt.Printf("Type: %T\n", intVar)
	fmt.Printf("Type: %T\n", floatVar)
	fmt.Printf("Type: %T\n", strVar)
	fmt.Printf("Type: %T\n", boolVar)
	fmt.Printf("Type: %T\n", pointerVar)

	fmt.Printf("Hexadecimal: %x\n", intVar)
	fmt.Printf("Hexadecimal: %x\n", strVar)
	fmt.Printf("Integer: %d\n", intVar)
	fmt.Printf("Float: %f\n", floatVar)

	fmt.Printf("Scientific (lowercase): %e\n", floatVar)
	fmt.Printf("Scientific (uppercase): %E\n", floatVar)

	fmt.Printf("String: %s\n", strVar)
	fmt.Printf("Pointer address: %p\n", pointerVar)
}
