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
		RequestProcessor:engine.Worker,

	}
	/*esimple:=engine.ConcurrentEngine{
		Scheduler:&scheduler.SimpleScheduler{},
		WorkerCount:5,
	}*/
	//Requests:=engine.Request{Url:seedUrl,Parser:engine.NewFuncParser(parser.ParseCityList,"ParseCityList")}
	Requests:=engine.Request{Url:"http://www.zhenai.com/zhenghun/shenzhen",Parser:engine.NewFuncParser(parser.ParseCity,"ParseCity")}
	e.Run(Requests)
}



