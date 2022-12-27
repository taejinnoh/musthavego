package main

import "fmt"

func main() {
	array := []int{1, 2, 3, 4, 5}
	slice2 := array[1:2]

	fmt.Println("array: ", array)
	fmt.Println("slice2: ", slice2, len(slice2), cap(slice2))

	array[1] = 100
	fmt.Println("After change second element")
	fmt.Println("array: ", array)
	fmt.Println("slice2: ", slice2, len(slice2), cap(slice2))

	slice2 = append(slice2, 500)
	fmt.Println("After append 500")
	fmt.Println("array: ", array)
	fmt.Println("slice2: ", slice2, len(slice2), cap(slice2))

}
