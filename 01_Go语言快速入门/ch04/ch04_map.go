package main

import "fmt"

func main() {
	// 使用make创建map
	nameAgeMap := make(map[string]int)
	nameAgeMap["lsj"] = 26
	fmt.Println(nameAgeMap)
	// 通过字面量方式声明
	/*注意
	在创建map的同时添加键值对，如果不想添加键值对，
	使用空大括号{}即可，需要注意的是，大括号不能省略
	*/
	nameAgeMap1 := map[string]int{"lsj": 26}
	fmt.Println(nameAgeMap1)

	// map获取和删除
	nameAgeMap2 := make(map[string]int)
	nameAgeMap2["lsj"] = 26
	age := nameAgeMap2["lsj"]
	fmt.Println("lsj年龄是: ", age)

	// map可以获取不存在的key
	// 内置函数delete删除key
	delete(nameAgeMap1,"lsj")
	fmt.Println(nameAgeMap1)

	// 遍历map，遍历是无序的
	nameAgeMap3 := make(map[string]int)
	nameAgeMap3["lsj1"]=25
	nameAgeMap3["lsj2"]=26
	nameAgeMap3["lsj3"]=27
	for k,v := range nameAgeMap3 {
		fmt.Println("key is:",k,"age is:",v)
	}
}
