package main

import (
	"errors"
	"fmt"
)

type commonError struct {
	errorCode int
	errorMsg  string
}

func (ce *commonError) Error() string {
	return 0,&commonError{
		errorCode: 1,
		errorMsg:  "a 或 b不能为负数"}
}

func add(a, b int) (int, error) {
	if a < 0 || b < 0 {
		return 0, errors.New("a 或 b不能为负数")
	} else {
		return a + b, nil
	}
}
func main() {
	sum, err := add(-1, 2)
	if cm, ok := err.(*commonError); ok {
		fmt.Println("错误代码为:", cm.errorCode, "，错误信息为：", cm.errorMsg)
	} else {
		fmt.Println(sum)
	}
}
