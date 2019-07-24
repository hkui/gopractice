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
		Url:"http://www.zhenai.com/zhenghun/beijing",
		Parser:worker.SerializedParser{
			Name:config.ParseCity,
			Args:"http://www.zhenai.com/zhenghun/beijing",
		},

	}

	var result worker.ParseResult
	err=client.Call(config.CrawlServiceRpc,req,&result)
	if err!=nil{
		t.Error(err)
	}else{
		for k,v:=range result.Items{
			fmt.Println(k,v.(map[string]interface{})["Url"])
		}

	}



}










