/*
1.结构体作为接收者
2.数组
3.类型和作用在它上面定义的方法必须在同一个包里定义，
	这就是为什么不能在 int、float 或类似这些的类型上定义方法。试图在 int 类型上定义方法会得到一个编译错误
	可通过预定义类型来搞
4.函数和方法的区别
	函数将变量作为参数：Function1(recv)

	方法在变量上被调用：recv.Method1()

 */
package main

import (
	"fmt"
	"time"
)

func main()  {
	stwo:=TwoNum{3,5}
	fmt.Println(stwo.Add())
	fmt.Println(stwo.Addanum(100))

	ar:=make(arr,5)
	ar[0]=1
	ar[1]=3
	ar[2]=5
	fmt.Println(ar.sum())//9

	m:=myTime{time.Now()}
	fmt.Println(m.String())
	fmt.Println(m.date())




}
//结构体
type TwoNum struct {
	a ,b int
}

func (tn *TwoNum)Add()int  {
	return tn.a+tn.b
}
func (tn *TwoNum)Addanum(num int)int  {
	return  tn.a+tn.b+num
}
//数组
type arr [] int

func (v arr)sum()(s int)  {
	for _,x:=range v{
		s+=x
	}
	return
}
//基本类型

type myTime struct {
	time.Time //anonymous field
}

func (t myTime)date() string  {
	return t.Time.String()[:10]

}
