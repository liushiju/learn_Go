package main

import "fmt"

func main(){
    array := [5]string{"a","b","c","d","e"}
    for i,v:= range array{
        fmt.Printf("索引:%d, 数值:%v\n",i ,v)
    }
}
