package main

import "fmt"

func swap(x, y *int) {
	*x, *y = *y, *x
}

func main() {
	a, b := 5, 10
	fmt.Println("beforeswap : a=", a, "b=", b)

	swap(&a, &b)
	fmt.Println("afterswap: a=", a, "b=", b)
}
