package main

import (
	"fmt"
	"imooc/retriever/mock"
)

type Retriever interface {
	Get(url string) string
}

func download(r Retriever)string  {
	return r.Get("http://www.baidu.com")
}

func main()  {
	var r Retriever
	r=mock.Retriever{"this is baidu"}
	fmt.Println(download(r))






}
