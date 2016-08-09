package main

//单向channel变量的声明非常简单，如下：
var ch1 chan int  // ch1是一个正常的channel，不是单向的
var ch2 chan<- float64// ch2是单向channel，只用于写float64数据
var ch3 <-chan int  // ch3是单向channel，只用于读取int数据

func main() {
	ch4 := make(chan int)
	ch5 := <-chan int(ch4) // ch5就是一个单向的读取channel
	ch6 := chan <- int(ch4) // ch6 是一个单向的写入channel
	print(ch4 ,ch5, ch6)

//	close channel
	close(ch4)
	close(ch5)
	close(ch6)

//	if not not closed
	x, ok := <-ch4
	println(x,"---", ok)
}
