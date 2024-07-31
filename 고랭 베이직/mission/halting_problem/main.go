package main

import (
	"fmt"
)

func halts(program func(), input int) bool {
	return false
}

func problematicFunction() {
	for {
		fmt.Println("this function never halts")
	}
}

func main() {

	problematicFunction()
	input := 42
	if halts(problematicFunction, input) {
		fmt.Println("the program halts")
	} else {
		fmt.Println("the program does not halt")
	}

}
