package main

import "C"
import "fmt"

/*
#include <stdlib.h>
*/

func Random() int {
	return int(C.random())
}

func Seed(i int) {
	C.srandom(C.uint(i))
}

func main() {
	Seed(1)
	fmt.Println(Random())
}
