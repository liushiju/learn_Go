package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["寒梅清风"] = 20
	fmt.Println("寒梅清风的年龄为",m["寒梅清风"])
	modifyMap(m)
	fmt.Println("寒梅清风的年龄为",m["寒梅清风"])
	
}


func modifyMap(p map[string]int)  {
	p["寒梅清风"]=22
}