package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5, 6}

	idx := 2

	//slice = append(slice, 0)

	//for i := len(slice) - 2; i >= idx; i-- {
	//	slice[i+1] = slice[i]
	//}
	//slice[idx] = 100
	//fmt.Println(slice)

	// 한 줄로 요소 삽입
	// 임시 슬라이스 {100,3,4,5,6}가 사용됨
	slice2 := append(slice[:idx], append([]int{100}, slice[idx:]...)...)
	fmt.Println(slice2)

	// 메모리 사용 없이 요소 삽입
	slice3 := append(slice, 0)
	copy(slice3[idx+1:], slice3[idx:]) // 슬라이싱으로 나온 슬라이스는 배열을 가리키는 포인터
	slice3[idx] = 100
	fmt.Println(slice3)

}
