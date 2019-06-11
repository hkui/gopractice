package engine

import (
	"fmt"
	"log"
)

type ConcurrentEngine struct {
	Scheduler Scheduler
	WorkerCount int
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
		fmt.Println("[get from out")
		result:=<-out
		fmt.Println("get from out]")

		for _,item:=range result.Items{
			itemCount++
			log.Printf("got item #%d %v",itemCount,item)
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
			fmt.Printf("send s.workerChan %v\n",in)
			notifier.WorkerReady(in)
			fmt.Println("[in->request")
			request:=<-in
			fmt.Println("in->request]")
			result,err:=worker(request)
			if err!=nil{
				continue
			}
			fmt.Println("[request->out")
			out<-result
			fmt.Println("request->out]")
		}
	}()
}
