package main

import (
	"fmt"
)

func main() {
	fmt.Println("Enter a number : ")
	var x int
	fmt.Scanf("%d", &x)

	if x < 0 {
		fmt.Println("The number is negative")
	} else if x == 0 {
		fmt.Println("The number is zero")
	} else if x > 0 && x <= 10 {
		fmt.Println("The number is between 1 and 10")
	} else {
		fmt.Println("The number is greater than 10")
	}
}
