package main

/*
结构体
*/
type Rect struct {
	x, y          float64
	width, height float64
}

func (r *Rect) Area() float64 {
	return r.height * r.width
}

/*
在Go语言中没有构造函数的概念，对象的创建通常交由一个全局的创建函数来完成，以
NewXXX 来命名，表示“构造函数”：
*/
func NewRect(x, y, width, height float64) *Rect {
	return &Rect{x, y, width, height}
}
func main() {
	rect1 := new(Rect)
	rect2 := &Rect{}
	rect3 := &Rect{0, 0, 100, 200}
	rect4 := &Rect{width: 100, height: 200}
	println(rect1.height)
	println(rect2.height)
	println(rect3.height)
	println(rect4.width)

}
