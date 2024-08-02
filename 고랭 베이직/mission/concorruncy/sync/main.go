package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() // 함수가 종료될 때, wg.done을 호출하여 waitGroup의 카운트를 줄인다.
	// 이는 현재 고루틴이 종료되었음을 나타내는 역하릉ㄹ 함
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second)
	fmt.Printf("worker %d done\n", id)
}

func main() {
	var wg sync.WaitGroup // waitGroup 인스턴스 형성. 모든 고 루틴이 완료될 때까지 메인 고루틴이 종료되지 않도록 한다

	for i := 1; i <= 3; i++ {
		wg.Add(1)         // waitGroup의 카운트를 1 증가시켜서 새로운 고루틴이 시작되었음을 나타냄
		go worker(i, &wg) // worker를 고루틴으로 실행.각 고루틴은 wg 포인터를 인자로 받아 작업이 끝날 때 wg.Done호출
	}

	wg.Wait()
} // 런타임 스케줄러에 의해 실제 실행되는 순서가 결정되기 때문에 계속 달라짐
