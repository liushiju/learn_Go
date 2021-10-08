package main

import (
	"errors"
	"fmt"
)

func sum(a int, b int) int {
	return a + b
}

// 多值返回
func sum2(c int, d int) (int, error) {
	if c < 0 || d < 0 {
		return 0, errors.New("c和d不能是负数")
	}
	return c + d, nil
}
func main() {
	result := sum(1, 2)
	fmt.Println(result)

	result2, err := sum2(6, 2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result2)
	}
}
