package main

import (
	"fmt"
	"net"
	"net/http"
	"net/rpc"
	"socket/main/server/library"
)

func main() {
	watcher := new(library.Watcher)
	rpc.Register(watcher)
	rpc.HandleHTTP()
	l, err := net.Listen("tcp", ":1234")
	if err != nil {
		fmt.Println("监听失败，端口可能已经被占用")
	}
	fmt.Println("正在监听1234端口")
	http.Serve(l, nil)

}
