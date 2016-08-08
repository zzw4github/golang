- 需要新建文件src
- intellij 最新的intellij才能装最新的插件 ，go插件需要最新的插件才能用最新的go 
- cgo windows 下安装 mingw 并将bin放在path下
- cgo.go stdlib 下 srandom random 改成 rand srand
- package main 下的main函数才是主函数 args下main.go 原来是 package args  go文件不能运行，改成 package main后可以。
- debug 有的go能debug，有的不能

关于gopath的设置之类的文章很多，看官自行go. 这里要在另外一个角度去解说gopath. 在我们以前熟悉的各种语言中都有这样几个概念： 系统路径,官方包路径,第三方包路径,项目路径。

好了go中只有两个路径.

GOROOT: go的安装路径,官方包路径根据这个设置自动匹配

GOPATH: 工作路径(其实不应该用中文翻译解释，直接说GOPATH更合适)

问题：项目路径和第三方包路径呢？ 首先：go中是没有项目这个概念的，只有包。可执行包只是特殊的一种，类似我们常说的项目 GOPATH可以设置多个，不管是可执行包，还是非可执行包，通通都应该在某个 $GOPATH/src下。

如果你这样做了，那么就不会出现本地包这种写法了

<!-- lang: cpp -->
import "./pathtopackage"
比如你可以把你的可执行(项目)包，安放在某个 $GOPATH/src下，例如 $GOPATH/src/app/youpackagedir

这样本地包的import就变成

<!-- lang: cpp -->
import "app/yourpackagedir/subpackage"
这样有什么用呢?

可以使用 go install 你的子包，有利于go build的时间，如果子包较大，那就更明显了
go code的自动完成可以用了