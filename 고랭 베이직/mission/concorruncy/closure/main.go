package main

import (
	"fmt"
)

// 클로저 : 함수와 그 함수가 정의된 컨택스트를 함께 저장한 객체이다
func main() {
	done := make(chan bool)

	for i := 0; i < 5; i++ {
		go func(num int) {
			fmt.Printf("%d", num)
			done <- true
		}(i) // 마지막 (i)는 즉시 실행이라는 의미를 가져, i의 현재값이 num으로 전달됨. 따라서 각 반복에서 다른 i값을 캡쳐하여 고루틴이 올바른 값을 갖느다
	}

	for i := 0; i < 5; i++ {
		<-done
	}

	// 비동기 환경에서 값의 변동을 피하기 위해서 i값을 복사하여 전달하는 것이 중요
}
