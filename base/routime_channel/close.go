/*
close函数是一个内建函数， 用来关闭channel，这个channel要么是双向的， 要么是只写的（chan<- Type）。
这个方法应该只由发送者调用， 而不是接收者。
当最后一个发送的值都被接收者从关闭的channel(下简称为c)中接收时,
接下来所有接收的值都会非阻塞直接成功，返回channel元素的零值。
如下的代码：
如果c已经关闭（c中所有值都被接收）， x, ok := <- c， 读取ok将会得到false。
*/
package main

import "fmt"

func main()  {
	ch:=make(chan  int,5)

	for i:=0;i<5;i++{
		ch<-i
	}
	//close(ch)

	for i:=0;i<10;i++{
		e,ok:=<-ch
		fmt.Printf("%v,%v\n",e,ok)
		if(!ok){
			break
		}
	}

}



