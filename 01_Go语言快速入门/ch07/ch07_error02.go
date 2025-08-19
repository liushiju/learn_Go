package main

import (
	"errors"
	"fmt"
)

func add(a, b int) (int, error) {
	if a < 0 || b < 0 {
		return 0, errors.New("a 或 b不能为负数")
	} else {
		return a + b, nil
	}
}

func main() {
	sum, err := add(-1, 2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(sum)
	}
}
