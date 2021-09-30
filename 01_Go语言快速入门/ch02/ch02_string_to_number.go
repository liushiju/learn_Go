package main

import (
	"fmt"
	"strconv"
)

func main() {
	i := 10
	i2s := strconv.Itoa(i)
	s2i, err := strconv.Atoi(i2s)
	fmt.Println(i2s, s2i, err)
}
