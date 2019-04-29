package main

import "fmt"

func numGen(start ,count int,out chan <-int)  {
	for i:=0;i<count ;i++  {
		out<-start
		start=start+count
	}
	close(out)
}
//close  https://blog.csdn.net/butterfly5211314/article/details/81842519
func numEchoRange(in<-chan int,done chan<- bool)  {
	for num:=range in  {
		fmt.Printf("%d\n",num)
	}
	done<-true
}

func main()  {
	numChan:=make(chan  int)
	done:=make(chan bool)
	go numGen(0,10,numChan)
	go numEchoRange(numChan,done)

	<-done

}

