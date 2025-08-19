package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func main() {
	p := person{Name: "寒梅清风", Age: 30}
	// struct to json
	jsonB, err := json.Marshal(p)
	if err == nil {
		fmt.Println(string(jsonB))
	}
	// json to struct
	respJSON := "{\"name\":\"李四\",\"age\":40}"
	json.Unmarshal([]byte(respJSON), &p)
	fmt.Println("Name is:", p.Name, "Age is:", p.Age)
	pt:= reflect.TypeOf(p)
	//遍历person字段中key为json的tag
	for i:=0;i<pt.NumField();i++{
		sf:=pt.Field(i)
		fmt.Printf("字段%s上,json tag为%s\n",sf.Name,sf.Tag.Get("json"))
	} 
}

type person struct {
	Name string `json:"name"`
	Age  uint   `json:"age"`
}
