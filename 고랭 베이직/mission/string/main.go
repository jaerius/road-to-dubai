package main

import (
	"fmt"
)

func main() {
	var str string = "Hello, Cosmos"
	fmt.Println(str)
	fmt.Println(len(str))
	fmt.Println(str[0])

	for index, runeValue := range str {
		fmt.Printf("%d: %c\n", index, runeValue)
	}

	str1 := "Hello, "
	str2 := "World!"
	combined := str1 + str2
	fmt.Println(combined)

	substr := str[7:13]
	fmt.Println(substr)

	byteSlice := []byte(str)
	byteSlice[0] = 'h'
	newStr := string(byteSlice)
	fmt.Println(newStr)
}
