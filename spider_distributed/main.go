package main

import (
	"spider/engine"
	"spider/scheduler"
	"spider/zhenai/parser"
	"spider_distributed/presist/client"
)

const seedUrl ="https://www.zhenai.com/zhenghun"
func main() {
	itemChan, err := client.ItemSaver(":1234")
	if err!=nil{
		panic(err)
	}

	e:=engine.ConcurrentEngine{
		Scheduler:&scheduler.QueuedScheduler{},
		WorkerCount:5,
		ItemChan:itemChan,

	}

	//Requests:=engine.Request{Url:seedUrl,ParseFunc:parser.ParseCityList}
	Requests:=engine.Request{Url:"http://www.zhenai.com/zhenghun/beijing",ParseFunc:parser.ParseCity}
	e.Run(Requests)
}



