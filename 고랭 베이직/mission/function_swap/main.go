package main

import (
	"fmt"
)

func swap(x, y string) (string, string) {
	return y, x
}

func namedReturnSwap(x, y string) (r1 string, r2 string) {
	r1, r2 = y, x
	return
}

func main() {
	x, y := "a", "b"
	cx, cy := swap(x, y)
	fmt.Println("x:", x, "y:", y)
	fmt.Println("cx:", cx, "cy:", cy)

	cx2, cy2 := namedReturnSwap(x, y)
	fmt.Println("cx2:", cx2, "cy2", cy2)
}
