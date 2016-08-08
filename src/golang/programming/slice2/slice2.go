package main

import "fmt"

func main() {
	mySlice := make([]int, 5, 10)
	mySlice = append(mySlice, 1, 2, 3, 4, 4, 4, 4, 44, 4, 4, 4, 4, 4, 4, 4, 4, 4)
	fmt.Println("len(mySlice):", len(mySlice))
	fmt.Println("cap(mySlice):", cap(mySlice))
}
