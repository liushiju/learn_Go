package main

import "testing"

func TestFibonacci(t *testing.T) {
	fsMap := map[int]int{}
	fsMap[0] = 0
	fsMap[1] = 1
	fsMap[2] = 2
	fsMap[3] = 3
	fsMap[4] = 4
	fsMap[5] = 5
	fsMap[8] = 8
	fsMap[13] = 13
	fsMap[21] = 21
	fsMap[34] = 34
	for k, v := range fsMap {
		fib := Fibonacci(k)
		if v == fib {
			t.Logf("结果正确:n为%d,值为%d", k, fib)
		} else {
			t.Errorf("结果错误：期望%d,但是计算的值是%d", v, fib)
		}
	}
}
