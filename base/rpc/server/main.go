package main

import (
	"base/rpc"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	rpc.Register(rpcdemo.DemoService{})
	listener, e := net.Listen("tcp", ":1234")
	if e!=nil{
		panic(e)
	}
	for{
		conn, e := listener.Accept()
		if e!=nil{
			log.Printf("accept error:%v",e)
			continue
		}
		go jsonrpc.ServeConn(conn)

	}
}
