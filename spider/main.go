package main

import (
	"spider/conf"
	"spider/engine"
	"spider/persist"
	"spider/scheduler"
	"spider/zhenai/parser"
)

const seedUrl ="https://www.zhenai.com/zhenghun"
func main() {
	itemChan, err := persist.ItemSaver(conf.EsConf)
	if err!=nil{
		panic(err)
	}

	e:=engine.ConcurrentEngine{
		Scheduler:&scheduler.QueuedScheduler{},
		WorkerCount:5,
		ItemChan:itemChan,

	}


	/*esimple:=engine.ConcurrentEngine{
		Scheduler:&scheduler.SimpleScheduler{},
		WorkerCount:5,
	}*/
	Requests:=engine.Request{Url:seedUrl,ParseFunc:parser.ParseCityList}
	//Requests:=engine.Request{Url:"http://www.zhenai.com/zhenghun/shenzhen",ParseFunc:parser.ParseCity}
	e.Run(Requests)
}



