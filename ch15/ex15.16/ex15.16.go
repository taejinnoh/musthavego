package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	var str string = "Hello World"
	var slice []byte = []byte(str)

	stringHeader1 := (*reflect.StringHeader)(unsafe.Pointer(&str))
	stringHeader2 := (*reflect.StringHeader)(unsafe.Pointer(&slice))

	fmt.Printf("Str:\t%x\n", stringHeader1)
	fmt.Printf("Str:\t%x\n", stringHeader2)
}
