package main

import "fmt"

func Counta(ch chan int) {
	ch <- 1
	fmt.Println("Counting")
}
func main() {
	chs := make([]chan int, 10)
	for i := 0; i < 10; i++ {
		chs[i] = make(chan int)
		go Counta(chs[i])
	}
	for _, ch := range chs {
		value := <-ch
		fmt.Println(value)
	}
}
