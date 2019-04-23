package main

import "fmt"

type person struct{
	name string
	sex int
	age byte
}

func main()  {
	p1:=new(person)  //&{ 0 0}
	p1.name="john"
	p1.age=10
	p1.sex=1
	fmt.Println(p1)

	var p2 person  //{ 0 0}
	var p3 *person  //<nil>
	p2.name="Jake"
	//p3.name="Lucy"

	fmt.Println(p2,p3)

}


