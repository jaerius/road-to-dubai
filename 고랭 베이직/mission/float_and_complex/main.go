package main

import (
	"fmt"
)

func main() {
	var f1 float32 = 3.14
	var f2 float64 = 3.141592653589793

	var c1 complex64 = complex(1.5, 2.5)
	var c2 complex128 = complex(1.5, 2.5)

	fmt.Println(f1)
	fmt.Println(f2)

	fmt.Println(real(c1))
	fmt.Println(imag(c1))

	fmt.Println(real(c2))
	fmt.Println(imag(c2))

	var c3 complex128 = complex(2.5, 3.5)
	var c4 complex128 = complex(1.5, -1.5)

	fmt.Println(c3 + c4)
	fmt.Println(c3 - c4)
	fmt.Println(c3 * c4)
	fmt.Println(c3 / c4)
}
