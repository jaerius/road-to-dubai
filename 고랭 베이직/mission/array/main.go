package main

import (
	"fmt"
)

func main() {
	var arr1 [5]int
	fmt.Println(arr1)

	arr2 := [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr2)

	arr3 := [5]int{1, 2}
	fmt.Println(arr3)

	fmt.Println(arr2[0])
	arr2[0] = 10
	fmt.Println(arr2[0])
}
