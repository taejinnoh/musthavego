package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func SumAtoB(a, b int) {
	sum := 0
	for i := a; i <= b; i++ {
		sum += i
	}
	fmt.Printf("%d부터 %d까지 합계는 %d입니다.\n", a, b, sum)
	wg.Done() // 작업 개수 감수
}

func main() {
	wg.Add(10) // 총 작업 개수 설정
	for i := 0; i < 10; i++ {
		go SumAtoB(i, 1000000000)
	}
	wg.Wait() // 작업 대기
}
