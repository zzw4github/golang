package main

import "sync"

var a string
var once sync.Once

func setup() {
	a = "hello, world"
}
func doprint() {
	once.Do(setup)
	print(a)
}
func twoprint() {
	go doprint()
	go doprint()
}

func main() {
	doprint()
	//twoprint()

}