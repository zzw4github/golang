package server

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

func main() {
	arith := new(Arith)
	rpc.Register(arith)
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	go http.Serve(l, nil)
	serverAddress := "http://192.168.0.1"
	client, err := rpc.DialHTTP("tcp", serverAddress+":8070")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	//	args := &server.Args{7, 8}
	args := &Args{7, 8}
	var reply int
	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Printf("Arith: %d*%d=%d", args.A, args.B, reply)
	//	quotient := new(Quotient)
	//	divCall := client.Go("Arith.Divide", args, &quotient, nil)
	//	replyCall := <-divCall.Done
}
