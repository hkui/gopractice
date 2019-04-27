package main

import (
	"fmt"
)

type any interface {

}
var i int =5
var str="ABC"

type Person1 struct {
	name string
	age int
}

func main()  {
	var val any
	val=5
	fmt.Println(val)

	val =str
	fmt.Println(val)

	per1:=new (Person1)
	per1.name="Rob"
	per1.age=10

	val=per1

	fmt.Println(val)//&{Rob 10}
	switch t:=val.(type){
	case int:
		fmt.Printf("%T\n",t)
	case string:
		fmt.Printf("%T\n",t)
	case bool:
		fmt.Printf("%T\n",t)
	case *Person1:
		fmt.Printf("%T\n",t)//*main.Person1
	default:
		fmt.Printf("%T\n",t)
	}



}

