package main

import (
	"fmt"
	"spider_distributed/config"
	"spider_distributed/rpcsupport"
	"spider_distributed/worker"
)

func main() {
	host:=fmt.Sprintf(":%d", config.WorkerPort0)
	err:=rpcsupport.ServeRpc(host, worker.CrawlService{})
	if err!=nil{
		panic(err)
	}
}
