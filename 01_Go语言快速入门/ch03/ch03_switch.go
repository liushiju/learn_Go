package main

import "fmt"

func main() {
	switch i := 6; {
	case i > 10:
		fmt.Println("i>10")
	case i>5 && i<=10:
		fmt.Println("i<5<10")
	default:
		fmt.Println("i<5")
	}
}