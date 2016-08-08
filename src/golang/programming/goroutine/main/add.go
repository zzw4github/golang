package main

import "fmt"

func Add(x, y int) {
	z := x + y
	fmt.Println(z)
}
func mainadd() {
	for i := 0; i < 10; i++ {
		go Add(i, i)
	}
}
