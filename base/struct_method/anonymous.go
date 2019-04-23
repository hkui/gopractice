/**
结构体匿名字段
嵌套结构体
赋值和初始化

 */
package main

import "fmt"

type student struct {
	 stuno int
	 grade string
	 int
	 Person

}
type Person struct {
	weight int
	height int
	sex int
}

func main()  {
	stu1:=new (student)
	stu1.stuno=11
	stu1.grade="0301"
	stu1.int=20

	stu1.weight=50
	stu1.height=165
	stu1.Person.height=180
	stu1.sex=1
	fmt.Println(stu1) //&{11 0301 20 {50 180 1}}

	stu2:=student{12,"032",33,Person{55,170,2}}
	fmt.Println(stu2)	//{12 032 33 {55 170 2}}



}
