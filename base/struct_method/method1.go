//指针和结构体作为接收者
package main

import "fmt"

type B struct {
	num int
}

func (b *B)change( v int )  {
	b.num=v
}
func (b B) wrie()string{
	return fmt.Sprint(b)
}
func main()  {
	var b1 B //b1是值
	b1.change(11)
	fmt.Println(b1.num) //11

	b2:=new(B) //b2是指针
	b2.change(20)
	fmt.Println(b2.num) //20
	fmt.Println(b1.num)  //11
	/**
	注意 Go 为我们做了探测工作，我们自己并没有指出是否在指针上调用方法，Go 替我们做了这些事情。
	b1 是值而 b2 是指针，方法都支持运行了
	 */
}

