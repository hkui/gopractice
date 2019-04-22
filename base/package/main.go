package main

import (
	"fmt"
	"./pack1"
	_ "./pack2" //执行init函数没被使用也不会报错
)

func main()  {

	fmt.Println(pack1.GetAddress())
	//fmt.Println(pack2.Salary)
}
//godoc -http=:6060 -goroot="."


