package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)

	wg.Add(1)
	go square(&wg, ch)
	for i := 0; i < 10; i++ {
		ch <- i * 2
	}
	close(ch) // 채널을 다 사용해서 닫음
	wg.Wait()
}

func square(wg *sync.WaitGroup, ch chan int) {
	for n := range ch { // 채널이 닫히면 종료
		time.Sleep(time.Second)
		fmt.Printf("Square: %d\n", n*n)
	}
	wg.Done()
}
