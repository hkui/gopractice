package main

import (
	"fmt"
	"runtime"
)
type worker struct {
	in chan int
	done chan  bool
}
func createWorker(id int)worker{
	w:=worker{
		in:make(chan int),
		done:make(chan  bool),
	}
	go dowork(id,w.in,w.done)
	return w
}

func dowork(id int,in chan int,done chan  bool )  {
	for n:=range in {
		fmt.Printf("worker %d reveived %c\n",id,n)
		go func() {
			done<-true
		}()

	}
}
const GNUM=3
func chanDemo()  {
	var workers [GNUM] worker
	for i:=0;i<GNUM;i++{
		workers[i]=createWorker(i)
	}
	fmt.Printf("Goroutine num:%d\n",runtime.NumGoroutine())

	for i,worker:=range workers{
		worker.in<-'a'+i
	}
	for i,worker:=range workers{
		worker.in<-'A'+i
	}
	fmt.Printf("Goroutine num:%d\n",runtime.NumGoroutine())
	for _,worker:=range  workers{
		<-worker.done
		<-worker.done
	}

}

func main() {
	chanDemo()

}
