package main

import (
	"fmt"
	"math/rand"
	"time"
)

func createWorker(id int) chan<- int {
	c := make(chan int)
	go worker(id, c)
	return c
}
func worker(id int, c chan int) {
	for n := range c {
		time.Sleep(2*time.Second)
		fmt.Printf("worker %d reveived %d\n", id, n)
	}
}
func generator() chan int {
	out := make(chan int)
	go func() {
		i := 0;
		for {
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond) //随机休眠
			out <- i
			i++
		}
	}()

	return out
}

func main() {
	c1, c2 := generator(), generator()
	w:=createWorker(0)
	var values []int
	n:=0
	hasValue:=false
	tm:=time.After(10*time.Second)
	tick:=time.Tick(time.Second)
	for {
		var activeWorker chan <-int  //nill channel
		activeValue:=0
		if hasValue{
			activeWorker=w
			activeValue=values[0]
		}

		select {
		case n= <-c1:
			{
				hasValue=true
				values=append(values, n)
			}
		case n= <-c2:
			{
				hasValue=true
				values=append(values, n)
			}
		case activeWorker<-activeValue:
			{
				hasValue=false
				values=values[1:]
			}
		case <-tick:{
			fmt.Printf("len(values)=%d\n",len(values))
		}
		case <-time.After(800*time.Millisecond):
			{
				fmt.Println("timeout")
			}
		case <-tm:
			{
				fmt.Printf("bye\n")
				return
			}


		}
	}

}
