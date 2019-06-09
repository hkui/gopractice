package main

import (
	"spider/engine"
	"spider/scheduler"
	"spider/zhenai/parser"
)

const seedUrl ="https://www.zhenai.com/zhenghun"
func main() {
	e:=engine.ConcurrentEngine{
		Scheduler:&scheduler.QueuedScheduler{},
		WorkerCount:5,

	}

	/*esimple:=engine.ConcurrentEngine{
		Scheduler:&scheduler.SimpleScheduler{},
		WorkerCount:5,
	}*/
	Requests:=engine.Request{Url:seedUrl,ParseFunc:parser.ParseCityList}
	e.Run(Requests)
}



