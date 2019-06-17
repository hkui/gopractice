package engine

import (
	"fmt"
)

type ConcurrentEngine struct {
	Scheduler Scheduler
	WorkerCount int
	ItemChan chan interface{}
}
type Scheduler interface {
	Submit( Request)
	WorkerChan() chan Request
	ReadyNotifier
	Run()
}
type ReadyNotifier interface {
	WorkerReady ( chan Request)
}

func (e *ConcurrentEngine)Run(seeds ...Request)  {
	out:=make(chan ParseResult)
	e.Scheduler.Run()

	for i:=0;i<e.WorkerCount;i++{
		createWorker(e.Scheduler.WorkerChan(),out,e.Scheduler)
	}
	for _,r:=range seeds{
		fmt.Println("send s.requestChan")
		e.Scheduler.Submit(r)
	}


	itemCount:=0
	for {
		result:=<-out

		for _,item:=range result.Items{
			itemCount++
			//log.Printf("got item #%d %v",itemCount,item)
			go func() {
				e.ItemChan <-item
			}()
		}
		for _,request:=range result.Requests{
			e.Scheduler.Submit(request)
		}
	}
}
//从Request通过fetcher和parser获取更多的Requests
func createWorker(in chan Request,out chan ParseResult,notifier ReadyNotifier)  {
	go func() {
		for{
			notifier.WorkerReady(in)
			request:=<-in
			result,err:=Worker(request)
			if err!=nil{
				continue
			}
			out<-result
		}
	}()
}
