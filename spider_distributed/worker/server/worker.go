package main

import (
	"flag"
	"fmt"
	"spider_distributed/rpcsupport"
	"spider_distributed/worker"
)
var port=flag.Int("port",0,"port to linten on")
//go run spider_distributed/worker/server/worker.go --port=9000
func main() {
	flag.Parse()
	if *port==0{
		fmt.Println("please specify port")
		return
	}
	host:=fmt.Sprintf(":%d", *port)
	err:=rpcsupport.ServeRpc(host, worker.CrawlService{})
	if err!=nil{
		panic(err)
	}
}
