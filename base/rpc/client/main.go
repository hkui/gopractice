package main

import (
	"base/rpc"
	"fmt"
	"net"
	"net/rpc/jsonrpc"
)

func main() {
	conn, e := net.Dial("tcp", ":1234")
	if e!=nil{
		panic(e)
	}
	client:=jsonrpc.NewClient(conn)

	var result float64
	e = client.Call("DemoService.Div", rpcdemo.Args{10, 2}, &result)
	if e!=nil{
		panic(e)
	}
	fmt.Println(result)


	e = client.Call("DemoService.Div", rpcdemo.Args{10, 0}, &result)
	if e!=nil{
		fmt.Println(e)
	}else {
		fmt.Println(result)
	}


}
