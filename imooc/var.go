package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"reflect"
	"runtime"
)

//欧拉公式
func euler(){
	//1i表示虚数,i不是
	fmt.Println(cmplx.Pow(math.E,1i*math.Pi)+1) //(0+1.2246467991473515e-16i),实部虚部均是float不准确
	cmplx.Exp(1i*math.Pi+1)
}
func plus(a,b int)int{
	return a+b;
}
func max(a,b int) int {
	if a>b {
		return a
	}
	return b
}

func apply(op func(a,b int)int ,a,b int)int  {
	p:=reflect.ValueOf(op).Pointer();
	opName:=runtime.FuncForPC(p).Name()
	fmt.Printf("calling func %s with args (%d,%d)\n",opName,a,b)
	return op(a,b)
}

func main()  {
	res:=apply(max,1,3)
	fmt.Printf("max=%d\n",res)
	//匿名函数方式
	apply(func(a, b int) int {
		return a*b
	},3,6)

}
