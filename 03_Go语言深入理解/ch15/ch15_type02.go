package main

import (
	"fmt"
	"reflect"
)

func main() {
	p := person{Name: "寒梅清风", Age: 30}
	pt := reflect.TypeOf(p)
	// 遍历person字段
	for i := 0; i < pt.NumField(); i++ {
		fmt.Println("字段:", pt.Field(i).Name)
	}
	// 遍历person的方法
	for i := 0; i < pt.NumMethod(); i++ {
		fmt.Println("方法:", pt.Method(i).Name)
	}
}

type person struct {
	Name string
	Age  uint
}

func (p person) String() string {
	return fmt.Sprintf("Name is %s,Age is %d", p.Name, p.Age)
}
