package main

import (
	"errors"
	"fmt"
)

func GetName() (firstName, lastName, nickName string) {
	return "first name", "last name", "nick name"
}

func Add(a int, b int) (ret int, err error) {
	if a < 0 || b < 0 { // 假设这个函数只支持两个非负数字的加法
		err = errors.New("Should be non-negative numbers!")
		return
	}
	return a + b, nil // 支持多重返回值
}

func main() {
	_, _, nickName := GetName()
	fmt.Println(nickName)
	ret, err := Add(1, 2)
	fmt.Println(err)
	fmt.Println(ret)
	if err != nil {
		fmt.Println("abc")
	}
}
