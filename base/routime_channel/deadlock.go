package main

import (
	"fmt"
	"time"
)

func f1(ch chan int)  {
	fmt.Println(<-ch)
}

func main()  {
	out:=make(chan int)
	out<-2
	go f1(out)
	time.Sleep(1e9)
}
//fatal error: all goroutines are asleep - deadlock!
