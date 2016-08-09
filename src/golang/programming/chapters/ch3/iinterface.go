package main

/*
在Go语言中，一个类只需要实现了接口要求的所有函数，我们就说这个类实现了该接口
*/

type File struct {
	// ...
}

func (f *File) Read(buf []byte) (n int, err error)                { return n, err }
func (f *File) Write(buf []byte) (n int, err error)               { return n, err }
func (f *File) Seek(off int64, whence int) (pos int64, err error) { return pos, err }
func (f *File) Close() error                                      { return nil }

type IFile interface {
	Read(buf []byte) (n int, err error)
	Write(buf []byte) (n int, err error)
	Seek(off int64, whence int) (pos int64, err error)
	Close() error
}
type IReader interface {
	Read(buf []byte) (n int, err error)
}
type IWriter interface {
	Write(buf []byte) (n int, err error)
}
type ICloser interface {
	Close() error
}

/*
尽管 File 类并没有从这些接口继承，甚至可以不知道这些接口的存在，但是 File 类实现了
这些接口，可以进行赋值：
*/
var file1 IFile = new(File)
var file2 IReader = new(File)
var file3 IWriter = new(File)
var file4 ICloser = new(File)

/*
Go语言的非侵入式接口，看似只是做了很小的文法调整，实则影响深远。
其一，Go语言的标准库，再也不需要绘制类库的继承树图。你一定见过不少C++、Java、C#
类库的继承树图。这里给个Java继承树图：
http://docs.oracle.com/javase/1.4.2/docs/api/overview-tree.html
在Go中，类的继承树并无意义，你只需要知道这个类实现了哪些方法，每个方法是啥含义
就足够了。
其二，实现类的时候，只需要关心自己应该提供哪些方法，不用再纠结接口需要拆得多细才
合理。接口由使用方按需定义，而不用事前规划。
其三，不用为了实现一个接口而导入一个包，因为多引用一个外部的包，就意味着更多的耦
合。接口由使用方按自身需求来定义，使用方无需关心是否有其他模块定义过类似的接口。
*/

//3.5.4 接口查询
//有办法让上面的 Writer 接口转换为 two.IStream 接口么？有。那就是我们即将讨论的接口
//查询语法，代码如下：
func main() {
	var file1 ICloser = new(File)
	if file5, ok := file1.(IFile); ok {
		println(file5)
	}
	//在Go语言中，你可以询问接口它指向的对象是否是某个类型
	if file6, ok := file1.(*File); ok {
		print(file6)
	}

	// type query

	var v1 interface{} = new(File)
	switch v := v1.(type) {
	case File:
		v.Close()
		print("v is type File") // 现在v的类型是int
	case string: // 现在v的类型是string

	}

	//type combination 组合
}
