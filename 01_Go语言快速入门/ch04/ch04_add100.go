package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i <= 100; i++ {
		if i%2 != 0 {
			continue
		}
		sum += i
	}
	fmt.Println("the sum is: ", sum)
}
