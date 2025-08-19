package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	s := "Hanson世纪"
	fmt.Printf("s的内存地址: %d\n", (*reflect.StringHeader)(unsafe.Pointer(&s)).Data)
	b := []byte(s)
	fmt.Printf("b的内存地址: %d\n", (*reflect.StringHeader)(unsafe.Pointer(&b)).Data)
	s3 := string(b)
	fmt.Printf("s3的内存地址: %d\n", (*reflect.StringHeader)(unsafe.Pointer(&s3)).Data)
}
