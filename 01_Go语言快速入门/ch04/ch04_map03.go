package main

import "fmt"

func main() {
	array := [3][2]string{{"a", "b"}, {"c", "d"}, {"e", "f"}}
	for i := range array {
		index := len(array[i])
		for j := 0; j < index; j++ {
			fmt.Println(array[i][j])
		}
	}
}
