package main

import "fmt"

// 입력값으로 Hello, 4를 입력한 경우 버퍼에 'ello 4\n'이 남게 되면서 연속적인 입력을 받을 수 없게 된다.
func main() {
	var a, b int
	n, err := fmt.Scan(&a, &b)
	if err != nil {
		fmt.Println(n, err)
	} else {
		fmt.Println(n, a, b)
	}
	n, err = fmt.Scan(&a, &b)
	if err != nil {
		fmt.Println(n, err)
	} else {
		fmt.Println(n, a, b)
	}
}
