package main

import (
	"fmt"
	"log"
	"os"
)

type Base struct {
	Name string
}

func (base *Base) Foo() {}

func (base *Base) Bar() {}

type Foo struct {
	Base
}

func (foo *Foo) Bar() {
	foo.Base.Bar()
}

/*
这段Go代码仍然有“派生”的效果，只是 Foo 创建实例的时候，需要外部提供一个 Base 类
实例的指针
匿名组合
*/
type Coo struct {
	*Base
}

type Job struct {
	Command     string
	*log.Logger //没有参数名
}

func NewJob(command string, logger *log.Logger) *Job {
	return &Job{command, logger}
}

func (job *Job) Start() {
	job.Logger.Println("starting now ...")
	job.Logger.Println("started...")
}

/*
显然这里会有问题。因为之前已经提到过，匿名组合类型相当于以其类型名称（去掉包名部分）
作为成员变量的名字。按此规则， Y 类型中就相当于存在两个名为 Logger 的成员，虽然类型不同。
因此，我们预期会收到编译错误。
有意思的是，这个编译错误并不是一定会发生的。假如这两个 Logger 在定义后再也没有被
用过，那么编译器将直接忽略掉这个冲突问题，直至开发者开始使用其中的某个 Logger 。
*/
type Logger struct {
	Level int
}
type Y struct {
	*Logger
	Name string
	//*log.Logger //duplicate field
}

func main() {
	job := &Job{"comman", log.New(os.Stderr, "Job:", log.Ldate)}
	job.Start()

}
