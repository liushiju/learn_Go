package main

import "fmt"

func main() {
	// 注意：[5]string和[4]string是不同类型
	// 长度也是数组类型的一部分
	array := [5]string{"a", "b", "c", "d", "e"}
	fmt.Println(array[2])
	// 长度可以忽略
	array1 := []string{"a", "b", "c", "d", "e"}
	fmt.Println(array1[2])
	// 针对指定元素初始化
	array2 := [5]string{1: "a", 3: "d"}
	fmt.Println(array2[1])
	// 通过for循环打印数组所有元素
	for i := 0; i < 5; i++ {
		fmt.Printf("数组索引 %d,对应值%s。", i, array[i])
	}
	// 返回值用不到，可以使用_下划线丢弃
	for _, v := range array {
		fmt.Printf("对应值%s\n", v)
	}
}
