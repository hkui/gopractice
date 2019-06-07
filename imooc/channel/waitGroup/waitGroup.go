package main

import (
	"fmt"
	"sync"
)
type worker struct {
	in   chan int
	Doff func()
}
func createWorker(id int,wg *sync.WaitGroup)worker{
	w:=worker{
		in:make(chan int),
		Doff: func() {
			wg.Done()
		},
	}
	go dowork(id,w)
	return w
}

func dowork(id int,w worker)  {
	for n:=range w.in {
		fmt.Printf("worker %d reveived %c\n",id,n)
		w.Doff();
	}
}
const GNUM=5
func chanDemo()  {
	var wg sync.WaitGroup

	var workers [GNUM] worker
	for i:=0;i<GNUM;i++{
		workers[i]=createWorker(i,&wg)
	}

	for i,worker:=range workers{
		wg.Add(1)
		worker.in<-'a'+i
	}
	for i,worker:=range workers{
		wg.Add(1)
		worker.in<-'A'+i
	}
	wg.Wait()

}

func main() {
	chanDemo()
}
