package main

import (
	"fmt"
	"imooc/retriever/mock"
	real2 "imooc/retriever/real"
)

type Retriever interface {
	Get(url string) string
}

func download(r Retriever)string  {
	return r.Get("http://www.baidu.com")
}

func main()  {
	var r Retriever
	r=mock.Mretriever{"this is baidu"}
	fmt.Println(download(r))

	fmt.Println()

	r2:=real2.Retriever{"oper",3}
	fmt.Println(download(r2))








}
