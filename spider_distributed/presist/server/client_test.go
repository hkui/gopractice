package main

import (
	"fmt"
	"spider/conf"
	"spider/model"
	"spider_distributed/rpcsupport"
	"testing"
	"time"
)
func TestSaver(t *testing.T) {
	const host=":12345"
	//启动server
	go ServeRpc(host,conf.EsConf)
	time.Sleep(2*time.Second)
	//启动client
	client, e := rpcsupport.NewClient(host)
	if e!=nil{
		panic(e)
	}
	item:=model.Profile{
			Id:1057746505,
			Url:"http://album.zhenai.com/u/1057746505",
			Name:"何处不相逢",
			Gender:"",
			Age:37 ,
			Height:158 ,
			Weight:0 ,
			Income:"12001-20000元",
			Marriage:"离异",
			Education:"硕士",
			Position:"北京",
		}

	result:=""
	//call
	e = client.Call("ItemSaverService.Save", item, &result)
	if e != nil || (result!="ok") {
		t.Errorf("result:%+v,err:%v", result, e)
	} else {
		fmt.Printf("result:%+v", result)
	}


}