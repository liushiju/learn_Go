package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["寒梅清风"] = 20
	fmt.Printf("main函数：m的内存地址为%p\n",m)
	fmt.Println("寒梅清风的年龄为",m["寒梅清风"])
	modifyMap(m)
	fmt.Println("寒梅清风的年龄为",m["寒梅清风"])
	
}


func modifyMap(p map[string]int)  {
	p["寒梅清风"]=22
	fmt.Printf("modifyMap函数：p的内存地址为%p\n",p)
}