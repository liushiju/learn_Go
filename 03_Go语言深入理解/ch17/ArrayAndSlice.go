package main

import "fmt"

func main() {
	ss := []string{"张三", "李四"}
	fmt.Println("切片ss长度为", len(ss), "容量为", cap(ss))
	ss = append(ss, "王五", "赵六")
	fmt.Println("切片ss长度为", len(ss), "容量为", cap(ss))
	fmt.Println(ss)
}
