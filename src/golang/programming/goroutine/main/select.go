package main

import "fmt"

func mainc() {
	cha := make(chan int, 1)
	for {
		select {
		case cha <- 0:
		case cha <- 1:
		}
		i := <-cha
		fmt.Println("Value received:", i)
	}
	//要创建一个带缓冲的channel，其实也非常容易：
	//	c := make(chan int, 1024)
	//	for i := range c {
	//		fmt.Println("Received:", i)
	//	}
	//	fatal error: all goroutines are asleep - deadlock!
}
