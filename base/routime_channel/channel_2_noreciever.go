package main

import "fmt"

func pump(ch chan int)  {
	for i:=0;;i++{
		ch<-i
	}
}

func main()  {
	ch:=make(chan int)
	go pump(ch)
	fmt.Println(<-ch)

	
}
