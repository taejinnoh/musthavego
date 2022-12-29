package main

import (
	"fmt"
	"sync"
	"time"
)

const goroutine = 10

type Job interface {
	Do()
}

type SquareJob struct {
	index int
}

func (j *SquareJob) Do() {
	fmt.Printf("%d 작업 시작\n", j.index)
	time.Sleep(1 * time.Second)
	fmt.Printf("%d 작업 완료 - 결과: %d\n", j.index, j.index*j.index)
}

func main() {
	var jobList [goroutine]Job
	for i := 0; i < goroutine; i++ {
		jobList[i] = &SquareJob{i}
	}
	var wg sync.WaitGroup
	wg.Add(goroutine)

	for i := 0; i < 10; i++ {
		job := jobList[i]
		go func() {
			job.Do()
			wg.Done()
		}()
	}
	wg.Wait()
}
