package main

import "fmt"

func Add(x, y int) {
	z := x + y
	fmt.Println(z)
}

/*
Go程序从初始化 main package 并执行 main() 函数开始，当 main() 函数返回时，程序退出，
且程序并不等待其他goroutine（非主goroutine）结束。
 */
func main() {
	for i := 0; i < 10; i++ {
		go Add(i, i)
	}
}
