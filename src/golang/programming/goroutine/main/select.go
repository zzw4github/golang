package main

import "fmt"
/*
Go语言直接在语言级别支持 select 关键字，用于处理异步IO
问题
 */
func main() {
	//cha := make(chan int, 1)
	//for {
	//	select {
	//	case cha <- 0:
	//	case cha <- 1:
	//	}
	//	i := <-cha
	//	fmt.Println("Value received:", i)
	//}

	/*
	要创建一个带缓冲的channel，其实也非常容易：创建了一个大小
	为1024的 int 类型 channel
	*/

	c := make(chan int, 1024)
	for   i :=0; i<1024;i++  {
		c <- 0
	}
	println("over")
	for i := range c {
		fmt.Println("Received:", i)
	}
	//	fatal error: all goroutines are asleep - deadlock!
}
