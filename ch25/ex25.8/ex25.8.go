package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	wg.Add(1)
	ctx, _ := context.WithTimeout(context.Background(), 3*time.Second) // 컨텍스트 생성
	go PrintEverySecond(ctx)
	//time.Sleep(5 * time.Second)
	//cancel() // 취소

	wg.Wait()
}

func PrintEverySecond(ctx context.Context) {
	tick := time.Tick(time.Second)
	for {
		select {
		case <-ctx.Done(): // 취소확인
			wg.Done()
			return
		case <-tick:
			fmt.Println("Tick")
		}
	}
}
