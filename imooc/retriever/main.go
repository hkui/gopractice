package main

import (
	"fmt"
	"imooc/retriever/mock"
	real2 "imooc/retriever/real"
	"time"
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
	//fmt.Println(download(r))
	//打出类型和值
	fmt.Printf("%T,%v\n",r,r)//mock.Mretriever,{this is baidu}
	inspect(r)

	rr:=r.(mock.Mretriever)
	fmt.Println("rr:",rr.Contents)




	fmt.Println()
	var r2 Retriever

	r2=&real2.Retriever{"oper",time.Minute}

	fmt.Printf("%T,%v\n",r2,r2)//*real.Retriever,&{oper 1m0s}

	//fmt.Println(download(r2))

	real_r2:=r2.(*real2.Retriever)// 需要 var r2 Retriever
	fmt.Println(real_r2.UserAgent)


	if mockRetriever,ok:=r.(mock.Mretriever);ok{
		fmt.Println(mockRetriever.Contents)
	}else{
		fmt.Println("not a mock retriever")
	}






}

func inspect(r Retriever)  {
	switch v:=r.(type) {
		case mock.Mretriever:{
			fmt.Println("inspect contents:",v.Contents)
		}
		case *real2.Retriever:{
			fmt.Println("inspect ua:",v.UserAgent)
		}
	}

}
