package main

import "fmt"

func main() {
	switch j := 1; j {
	case 1:
		fallthrough
	case 2:
		fmt.Println("2")
	default:
		fmt.Println("error")
	}
}