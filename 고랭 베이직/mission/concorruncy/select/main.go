package main

import (
	"fmt"
)

func main() {
	ch := make(chan int, 1) // 버퍼의 크기 1인, int를 보낼 수 있는 채널을 만들어라
	ch <- 1                 // 보내자 마자 채널 닫힘
	select {
	case val := <-ch: // ch채널에서 값을 수신하려고 시도, 값이 있으면 val에 저장
		fmt.Println("recevied", val)
	default:
		fmt.Println("no value recevied")
	}
}
