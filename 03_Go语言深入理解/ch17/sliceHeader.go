package main

import (
	"fmt"
	"unsafe"
)

func main() {
	a1 := [2]string{"张三", "李四"}
	s1 := a1[0:1]
	//s2 := a1[:]
	//	打印出s1和s2的Data值，是一样的
	//fmt.Println((*reflect.SliceHeader)(unsafe.Pointer(&s1)).Data)
	//fmt.Println((*reflect.SliceHeader)(unsafe.Pointer(&s2)).Data)

	sh1 := (*slice)(unsafe.Pointer(&s1))
	fmt.Println(sh1.Data, sh1.Len, sh1.Cap)
}

type slice struct {
	Data uintptr
	Len  int
	Cap  int
}
