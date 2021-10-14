package main

import (
	"fmt"
	"sync")


func main() {
	var person sync.Map
	// 将键值对保存在sync.Map
	person.Store("张三", 26)
	person.Store("李四", 30)
	person.Store("王五", 32)
	person.Store("赵六", 37)
	// 从sync.Map中根据键取值
	fmt.Println(person.Load("张三"))
	// 根据键删除对应的键值对
	person.Delete("张三")
	// 遍历所有的sync.Map中的键值对
	person.Range(func(key, value interface{}) bool {
		fmt.Println("iterate: ", key,value)
		return true
	})

}
