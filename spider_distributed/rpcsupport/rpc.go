package rpcsupport

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)
//注册并开一个服务
func ServeRpc(host string,service interface{}) error {
	err := rpc.Register(service)
	if err!=nil{
		panic(err)
	}
	listener, e := net.Listen("tcp", host)
	if e!=nil{
		return e
	}
	log.Printf("rpc-itemSaver-server waiting!\n")
	for{
		conn, e := listener.Accept()
		if e!=nil{
			log.Printf("accept error:%v",e)
			continue
		}
		go jsonrpc.ServeConn(conn)

	}
	return nil
}
//创建客户端
func NewClient(host string)(* rpc.Client,error)  {
	conn, e := net.Dial("tcp", host)
	if e!=nil{
		return nil,e
	}
	return jsonrpc.NewClient(conn),nil

}