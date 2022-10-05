package main

import "fmt"

func main() {
	temp := 18

	switch true {
	case temp < 10, temp > 18:
		fmt.Println("바깥 활동하기 좋은 날씨가 아닙니다.")
	case temp >= 10 && temp < 20:
		fmt.Println("약간 추울 수 있으니 가벼운 겉옷을 준비하세요.")
	default:
		fmt.Println("따뜻합니다.")
	}
}
