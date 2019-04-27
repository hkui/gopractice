package main

import (
	"fmt"
	"time"
)

func sendData(ch chan  string,str string)  {
	for i:=0;i<3 ;i++  {
		ch<-fmt.Sprintf("%s_%d",str,i)
		time.Sleep(1e9)
	}

}
func getData(ch chan string)  {
	for{
		fmt.Printf("%s",<-ch)
	}
}
func main()  {
	ch:=make(chan string)
	go sendData(ch,"php")
	go getData(ch)

	time.Sleep(5*1e9)


}
