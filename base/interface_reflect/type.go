/*
1.类型推断
2.type stwitch
3.测试一个值是否实现了某个接口
*/

package main

import "fmt"

type Animal interface {
	shot()
}
type dog struct {
	name string
}

func (r dog) shot() {
	fmt.Println(r.name + ":shot")
}

type Person interface {
	intro()string
}
type student struct {
	name string
	age int
}

func (r student)intro()string  {
	return fmt.Sprintf("name=%s,age=%d",r.name,r.age)
}

func main() {
	var a1, a2 Animal
	d1 := new(dog)
	d1.name = "poly"
	a1 = d1
	d2 := dog{"harry"}
	a2 = d2;
	if t, ok := a1.(*dog); ok {
		fmt.Printf("type of a1=%T\n", t) //type of a1=*main.dog
	}
	if t, ok := a2.(dog); ok {
		fmt.Printf("type of a2=%T\n", t) //type of a2=main.dog
	}

	/**
	所有 case 语句中列举的类型（nil 除外）都必须实现对应的接口，
	如果被检测类型没有在 case 语句列举的类型中，就会执行 default 语句。

	可以用 type-switch 进行运行时类型分析，但是在 type-switch 不允许有 fallthrough
	 */
	switch t := a1.(type) {
	case *dog:
		fmt.Printf("Type=%T", t)
	case nil:
		fmt.Printf("nil")
	default:
		fmt.Printf("Unexpected type %T\n", t)
	}
	fmt.Println("\n--------------------------\n")

	var stu Person=student{"hk",18}
	

	if sv,ok:=stu.(Person);ok{
		fmt.Printf("stu implements Person():%s\n",sv.intro())
	}



}
