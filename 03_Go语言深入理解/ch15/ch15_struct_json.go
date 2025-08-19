package main

import (
	"fmt"
	"reflect"
	"strings"
)

func main() {
	p := person{Name: "寒梅清风", Age: 20}
	pv := reflect.ValueOf(p)
	pt := reflect.TypeOf(p)
	//自己实现的struct to json
	jsonBuilder := strings.Builder{}
	jsonBuilder.WriteString("{")
	num := pt.NumField()
	for i := 0; i < num; i++ {
		jsonTag := pt.Field(i).Tag.Get("json") //获取json tag
		jsonBuilder.WriteString("\"" + jsonTag + "\"")
		jsonBuilder.WriteString(":")
		//获取字段的值
		jsonBuilder.WriteString(fmt.Sprintf("\"%v\"", pv.Field(i)))
		if i < num-1 {
			jsonBuilder.WriteString(",")
		}
	}
	jsonBuilder.WriteString("}")
	fmt.Println(jsonBuilder.String()) //打印json字符串
}

type person struct {
	Name string `json:"name"`
	Age  uint `json:"age"`
}
