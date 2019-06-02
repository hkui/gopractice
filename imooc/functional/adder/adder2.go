package main

import "fmt"

type iAdder func(int)(int,iAdder)

func adder2(base int) iAdder  {
	return func(v int) ( int, iAdder) {
		return base+v,adder2(base+v)
	}
}
func main()  {
	fun:=adder2(0)
	for i:=0;i<10;i++{
		var s int
		s,fun=fun(i)
		fmt.Printf("%d,%d\n",i,s)

	}
}
