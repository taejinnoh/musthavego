package main

import "fmt"

func main() {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := make([]int, 3, 10)
	slice3 := make([]int, 10)
	slice4 := append([]int{}, slice1...)

	cnt1 := copy(slice2, slice1)
	cnt2 := copy(slice3, slice1)

	slice1[2] = 100

	fmt.Println("slice2:", slice2, len(slice2), cap(slice2), "copy_cnt:", cnt1)

	fmt.Println("slice3:", slice3, len(slice3), cap(slice3), "copy_cnt:", cnt2)

	fmt.Println("slice4:", slice4, len(slice4), cap(slice4))
}
