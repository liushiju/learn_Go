package main

import "fmt"

func main() {
	sum := 0
	i := 1
	for {
		sum += i
		i++
		if i > 100 {
			break
		}

	}
	fmt.Println("the sum is:", sum)
}
