package main

import "fmt"

func main()  {
	m:=map[string]string{"a":"ajax","p":"php"}
	for k,v:=range m{
		fmt.Println(k,v)
	}
	m1:=make(map[string]string)
	var m2 map[string]string

	fmt.Println(m1,m2)
}
