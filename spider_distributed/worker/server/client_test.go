package main

import (
	"fmt"
	"spider_distributed/config"
	"spider_distributed/rpcsupport"
	"spider_distributed/worker"
	"testing"
	"time"
)

func TestCrawlService(t *testing.T)  {
	const host=":9000"
	go rpcsupport.ServeRpc(host,worker.CrawlService{})
	time.Sleep(time.Second)

	client,err:=rpcsupport.NewClient(host)
	if err!=nil{
		panic(err)
	}
	req:=worker.Request{
		Url:"http://album.zhenai.com/u/1875061106",
		Parser:worker.SerializedParser{
			Name:config.ParseProfile,
			Args:"http://album.zhenai.com/u/1875061106",
		},

	}
	var result worker.ParseResult
	err=client.Call(config.CrawlServiceRpc,req,&result)
	if err!=nil{
		t.Error(err)
	}else{
		fmt.Printf("r=%v\n",result)
	}



}










