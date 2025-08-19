package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	s := "张三"
	//b := []byte(s)
	sh := (*reflect.SliceHeader)(unsafe.Pointer(&s))
	sh.Cap = sh.Len
	b1 := *(*[]byte)(unsafe.Pointer(sh))
	fmt.Println(b1)
}
