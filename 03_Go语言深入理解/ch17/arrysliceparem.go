package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	a1 := [2]string{"张三", "李四"}
	fmt.Printf("函数main数组指针: %p\n", &a1)
	arrayF(a1)
	s1 := a1[0:1]
	fmt.Println((*reflect.SliceHeader)(unsafe.Pointer(&s1)).Data)
	sliceF(s1)
}

func sliceF(s []string) {
	fmt.Printf("函数slceF Data: %d\n", (*reflect.SliceHeader)(unsafe.Pointer(&s)).Data)
}

func arrayF(a [2]string) {
	fmt.Printf("函数arrayF数组指针: %p\n", &a)
}
