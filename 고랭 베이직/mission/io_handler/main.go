package main

import (
	"fmt"
)

func main() {

	for {
		var x string
		fmt.Println("Enter command: ")
		fmt.Scanf("%s", &x)

		switch x {
		case "exit":
			fmt.Println("Existing...")
			return
		case "hello":
			fmt.Println("hello, world!")
		case "even":
			fmt.Println("Even numbers from 0 to 10: ")
			fmt.Println("0 2 4 6 10")
		case "odd":
			fmt.Println("Odd numbers from 1 to 10: ")
			fmt.Println("1 3 5 7 9")
		default:
			fmt.Println("Unknown command")
		}
	}
}
