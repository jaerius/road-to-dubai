package main

import (
	"fmt"
)

func main() {
	var s1 []int
	fmt.Println(s1)

	s2 := []int{1, 2, 3, 4, 5}
	fmt.Println(s2)

	s3 := make([]int, 5)
	fmt.Println(s3)

	fmt.Println(s2[0])
	s2[0] = 10
	fmt.Println(s2[0])

	subSlice := s2[1:3]
	fmt.Println(subSlice)

	fmt.Println(len(s2))
	fmt.Println(cap(s2))

	s2 = append(s2, 6)
	fmt.Println(s2)
}
