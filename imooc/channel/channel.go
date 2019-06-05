package main

import (
	"fmt"
	"time"
)

func worker(id int,c chan int  )  {
	/*for{
		n,ok:=<-c
		if !ok{
			break
		}
		fmt.Printf("worker %d reveived %c\n",id,n)
	}*/

	for n:=range c {
		fmt.Printf("worker %d reveived %c\n",id,n)
	}
}

//返回的chan只能接收数据
// <-chan int 表示只能发送数据
func createWorker(id int)chan <-int {
	c:=make(chan int)
	go func() {
		for    {
			fmt.Printf("worker %d reveived %d\n",id,<-c)
		}
	}()

	return c

}

func chanDemo()  {
	var channels [10]chan <-int

	for i:=0;i<10;i++{
		channels[i]=createWorker(i)
	}
	for i:=0;i<10;i++{
		channels[i]<-i+'a'
	}
	for i:=0;i<10;i++{
		channels[i]<-i+'A'
	}


	time.Sleep(time.Millisecond)

}

func bufferedChannel()  {
	c:=make(chan int,3)
	go worker(0,c)
	c<-'a'
	c<-'b'
	c<-'c'

	time.Sleep(time.Millisecond)
}

func channelClose()  {
	c:=make(chan int,3)
	go worker(0,c)
	c<-'a'
	c<-'b'
	c<-'c'
	//发送方才能close,close后读方还会一直读，读到的是默认的初始空
	close(c)
	time.Sleep(time.Millisecond)

}

func main() {
	//chanDemo()
	//bufferedChannel()
	channelClose()
}
