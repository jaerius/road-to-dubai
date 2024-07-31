package main

import (
	"fmt"
)

func main() {
	var a int = 10
	var ptr *int

	ptr = &a // 포인터 변수는 a의 주소값을 가지고 있다

	fmt.Println(a)
	fmt.Println("Address of a:", ptr)
	fmt.Println("value at the address stored in ptr:", *ptr)

	*ptr = 20
	fmt.Println("new value of a:", a)

}
