package main

import "fmt"

const (
	s1 = 3 + iota
	s2 = 9
	s3
	s4
)

func main() {
	fmt.Printf("s1=%d,s2=%d,s3=%d,s4=%d\n",s1,s2,s3,s4)
}