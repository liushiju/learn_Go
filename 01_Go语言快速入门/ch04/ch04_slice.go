package main

import "fmt"

func main() {
	array := [5]string{"a", "b", "c", "d", "e"}
	// 索引包含2不包含5
	slice := array[2:5]
	fmt.Println(slice)
	// 切片修改
	/* 数组对应的值已经被修改为 f，所以这也证明了
	基于数组的切片，使用的底层数组还是原来的数组，
	一旦修改切片的元素值，那么底层数组对应的值也
	会被修改*/
	fmt.Println("修改切片前: ", array)
	slice1 := array[2:5]
	slice1[1] = "f"
	fmt.Println("修改切片后: ", array)
	// 切片声明
	// 元素类型string，长度4
	// slice2 := make([]string,4)
	// 元素类型string，长度4，容量8
	// 注意：切片的容量不能比切片的长度小
	// slice3 := make([]string,4,8)

	// 切片还可以通过字面量方式声明和初始化
	/*
		切片和数组的字面量初始化方式，差别就是中括号里
		的长度。此外，通过字面量初始化的切片，长度和容量
		相同
	*/
	slice4 := []string{"a", "b", "c", "d", "e"}
	fmt.Println(len(slice4), cap(slice4))
	// append追加切片元素
	// slice5:=append(slice4,"f")
	slice5 := append(slice4, "f", "g")
	fmt.Println(slice5)
	slice6 := append(slice1, slice5...)
	fmt.Println(slice6)
	/*注意
	在创建新切片的时候，最好要让新切片的长度和容量
	一样，这样在追加操作的时候就会生成新的底层数组，
	从而和原有数组分离，就不会因为共用底层数组导致
	修改内容的时候影响多个切片
	*/

	// 切片元素循环
	for i, v := range slice6 {
		fmt.Printf("切片索引%d,对应值:%s\n", i, v)
	}
}
