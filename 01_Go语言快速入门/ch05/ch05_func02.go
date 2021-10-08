package main

import (
	"errors"
	"fmt"
)

func sum(a, b int) (sum int, err error) {
	if a < 0 || b < 0 {
		return 0, errors.New("a或b不能是负数")
	}
	sum = a + b
	err = nil
	return
}

func main() {
	result, err := sum(10, 2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}
