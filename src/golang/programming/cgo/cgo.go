package main

/*
#include <stdlib.h>
*/
import "C" //  包含的头文件和import statement中间不能带空行

import "fmt"

func Random() int {
	return int(C.rand())
}
func Seed(i int) {
	C.srand(C.uint(i))
}
func main() {
	Seed(100)
	fmt.Println("Random:", Random())
}
