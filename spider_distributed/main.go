package main

import (
	"flag"
	"log"
	"net/rpc"
	"spider/engine"
	"spider/scheduler"
	"spider/zhenai/parser"
	"spider_distributed/config"
	itemsaver "spider_distributed/presist/client"
	"spider_distributed/rpcsupport"
	worker "spider_distributed/worker/client"
	"strings"
)

const seedUrl = "https://www.zhenai.com/zhenghun"

var (
	itemSaverHost=flag.String("itemsaver_host","","itemsaver_host")
	wokerHosts=flag.String("worker_hosts","","worker hosts(comma separated)")

)

func main() {
	flag.Parse()
	if *itemSaverHost==""{
		log.Printf("please specify itemSaverHost ")
		return
	}
	if *wokerHosts==""{
		log.Printf("please specify wokerHosts ")
		return
	}

	itemChan, err := itemsaver.ItemSaver(*itemSaverHost)
	if err != nil {
		panic(err)
	}
	pool:=createClientPool(strings.Split(*wokerHosts,","))
	processor,_:=worker.CreateProcessor(pool)


	e := engine.ConcurrentEngine{
		Scheduler:        &scheduler.QueuedScheduler{},
		WorkerCount:      5,
		ItemChan:         itemChan,
		RequestProcessor: processor,
	}

	//Requests:=engine.Request{Url:seedUrl,ParseFunc:parser.ParseCityList}
	Requests := engine.Request{Url: "http://www.zhenai.com/zhenghun/beijing",
		Parser: engine.NewFuncParser(
			parser.ParseCity, config.ParseCity,
		),}
	e.Run(Requests)
}
func createClientPool(hosts []string) chan *rpc.Client {
	var clients []*rpc.Client
	for _, h := range hosts {
		client, err := rpcsupport.NewClient(h)
		if err == nil {
			clients = append(clients, client)
			log.Printf("connected to %s", h)
		} else {
			log.Printf("Error connecting to %s:%v", h, err)
		}
	}
	out := make(chan *rpc.Client)

	go func() {
		for {
			for _, client := range clients {
				out <- client
			}
		}

	}()
	return out
}
