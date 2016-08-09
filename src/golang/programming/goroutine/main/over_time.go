package main

import "time"
import "fmt"

func main() {
	// 首先，我们实现并执行一个匿名的超时等待函数
	timeout := make(chan bool, 1)
	ch := make(chan bool, 1)
	go func() {
		time.Sleep(2e9) // 等待1秒钟
		ch <- false
	}()

	go func() {
		time.Sleep(1e9) // 等待1秒钟
		timeout <- true
	}()
	// 然后我们把timeout这个channel利用起来
	select {
	case value := <-ch:
		// 从ch中读取到数据
		fmt.Println(value)
	case value := <-timeout:
		// 一直没有从ch中读取到数据，但从timeout中读取到了数据
		fmt.Println(value)
	}
}
