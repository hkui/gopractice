package main

import (
	"spider/engine"
	"spider/persist"
	"spider/scheduler"
	"spider/zhenai/parser"
)

const seedUrl ="https://www.zhenai.com/zhenghun"
func main() {
	e:=engine.ConcurrentEngine{
		Scheduler:&scheduler.QueuedScheduler{},
		WorkerCount:5,
		ItemChan:persist.ItemSaver(),
	}


	/*esimple:=engine.ConcurrentEngine{
		Scheduler:&scheduler.SimpleScheduler{},
		WorkerCount:5,
	}*/
	//Requests:=engine.Request{Url:seedUrl,ParseFunc:parser.ParseCityList}
	Requests:=engine.Request{Url:"http://www.zhenai.com/zhenghun/beijing",ParseFunc:parser.ParseCity}
	e.Run(Requests)
}



