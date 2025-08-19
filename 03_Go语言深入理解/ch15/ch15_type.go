package main

import (
	"fmt"
	"reflect"
)

func main() {
	p := person{Name: "寒梅清风", Age: 30}
	ppv := reflect.ValueOf(&p)
	fmt.Println(ppv.Kind())
	pv := reflect.ValueOf(p)
	fmt.Println(pv.Kind())
}

type person struct {
	Name string
	Age  uint
}
