package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	p := person{Name: "寒梅清风", Age: 30}
	// struct to json
	jsonB, err := json.Marshal(p)
	if err == nil {
		fmt.Println(string(jsonB))
	}
	// json to struct
	respJSON := "{\"Name\":\"李四\",\"Age\":40}"
	json.Unmarshal([]byte(respJSON), &p)
	fmt.Println("Name is:",p.Name,"Age is:",p.Age)
}

type person struct {
	Name string
	Age  uint
}
