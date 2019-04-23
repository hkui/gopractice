package main

import (
	"fmt"
	"strconv"
)

type Twoint struct {
	a,b int
}

func (tn *Twoint)String()string  {
	return "(" + strconv.Itoa(tn.a) + "/" + strconv.Itoa(tn.b) + ")"
}
func main()  {
	two:=new(Twoint)
	two.a=3
	two.b=5
	fmt.Printf("two is:%v\n",two)
	fmt.Println("two =",two)
	fmt.Printf("two=%T\n",two)
	fmt.Printf("two=%#v\n",two)
	/**
two is:(3/5)
two = (3/5)
two=*main.Twoint
two=&main.Twoint{a:3, b:5}
	
当你广泛使用一个自定义类型时，最好为它定义 String()方法。
从上面的例子也可以看到，格式化描述符 %T 会给出类型的完全规格，
%#v 会给出实例的完整输出，包括它的字段（在程序自动生成 Go 代码时也很有用）
	 */

}
