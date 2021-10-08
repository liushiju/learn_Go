package main

import (
	"fmt"
	"sort"
)

func main() {
	nameAgeMap := make(map[string]int)
	nameAgeMap["张三1"] = 25
	nameAgeMap["张三2"] = 26
	nameAgeMap["张三3"] = 27
	nameAgeMap["张三4"] = 28
	nameAgeMap["张三5"] = 29

	array := []string{}
	for k := range nameAgeMap {
		array = append(array, k)
	}
	fmt.Println(array)
	sort.Strings(array)
	fmt.Println(array)
	for _, v := range array {
		fmt.Printf("%s: %d\n", v, nameAgeMap[v])
	}
	// map没有容量，只有长度，即大小
	fmt.Println("键值对长度: ", len(nameAgeMap))
	
}
