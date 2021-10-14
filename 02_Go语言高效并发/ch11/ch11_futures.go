package main

import (
	"fmt"
	"time")


func washVegetables() <-chan string {
	vegetables := make(chan string)
	go func() {
		time.Sleep(5*time.Second)
		vegetables <- "洗好的菜"
	}()
	return vegetables
}

func boilWater() <- chan string  {
	water := make(chan string)
	go func()  {
		time.Sleep(5*time.Second)
		water <- "烧开的水"
	}()
	return water
}

func main()  {
	vegetablesCh := washVegetables()
	waterCh := boilWater()
	fmt.Println("已经安排洗菜和烧水了，我先眯一会")
	time.Sleep(2*time.Second)

	fmt.Println("要做火锅了，看看菜和水好了吗")
	vegetables := <- vegetablesCh
	water := <- waterCh
	fmt.Println("准备好了，可以做火锅了:",vegetables,water)
}