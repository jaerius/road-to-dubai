package main

import (
	"fmt"
)

func send(ch chan<- int, val int) { // ch가 chan<- int 타입
	ch <- val
}

func receive(ch <-chan int) int { // ch가 <-chan int 타입
	return <-ch // 채널에서 값을 받아 반환하는 의미
}

func main() {
	ch := make(chan int)

	go send(ch, 42)
	val := receive(ch)
	fmt.Println(val)
}
