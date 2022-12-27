package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

const (
	Balance       = 1000 // 초기 잔액
	EarPoint      = 500  // 승리수당
	LosePoint     = 100  // 패배 수당
	VictoryPoint  = 5000 // 게임 승리 포인트
	GameoverPoint = 0    // 게임 패배 포인트
)

var stdin = bufio.NewReader(os.Stdin)

func InputValue() (int, error) {
	var n int
	_, err := fmt.Scanln(&n)
	if err != nil {
		stdin.ReadString('\n')
	}
	return n, err
}

func main() {
	balance := Balance
	rand.Seed(time.Now().UnixNano())

	for {
		fmt.Printf("1~5 사이 값을 입력하세요: ")
		value, err := InputValue()
		if err != nil {
			fmt.Println("숫자를 입력하세요.")
		} else if value > 5 || value < 1 {
			fmt.Println("1~5 사이 값만 입력하세요.")
		} else {
			r := rand.Intn(5) + 1
			if value != r {
				balance -= LosePoint
				fmt.Println("틀렸습니다. 잔액: ", balance)
				if balance <= GameoverPoint {
					fmt.Println("게임 오버")
					break
				}
			} else {
				balance += EarPoint
				fmt.Println("축하합니다. 잔액: ", balance)
				if balance >= VictoryPoint {
					fmt.Println("게임 승리")
					break
				}
			}
		}
	}
}
