package main

import (
	"spider/engine"
	"spider/scheduler"
	"spider/zhenai/parser"
	"spider_distributed/config"
	itemsaver "spider_distributed/presist/client"
	worker "spider_distributed/worker/client"
)

const seedUrl ="https://www.zhenai.com/zhenghun"
func main() {
	itemChan, err := itemsaver.ItemSaver(":1234")
	if err!=nil{
		panic(err)
	}
	processor,err:=worker.CreateProcessor()
	if err!=nil{
		panic(err)
	}

	e:=engine.ConcurrentEngine{
		Scheduler:&scheduler.QueuedScheduler{},
		WorkerCount:5,
		ItemChan:itemChan,
		RequestProcessor:processor,

	}

	//Requests:=engine.Request{Url:seedUrl,ParseFunc:parser.ParseCityList}
	Requests:=engine.Request{Url:"http://www.zhenai.com/zhenghun/beijing",
		Parser:engine.NewFuncParser(
			parser.ParseCity,config.ParseCity,
		),}
	e.Run(Requests)
}



