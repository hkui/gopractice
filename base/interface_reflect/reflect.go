/**
1.反射基本类型
2.反射结构体
 */
package main

import (
	"fmt"
	"reflect"
)
type unknowstruct struct {
	a int
	b string
	C float32
}

func (s unknowstruct)Intro() string {
	return fmt.Sprintf("a=%d,b=%s,c=%f",s.a,s.b,s.C)
}
func (s unknowstruct)private() string {
	return fmt.Sprintf("private a=%d,b=%s,c=%f",s.a,s.b,s.C)
}

var secret interface{}=unknowstruct{2,"php",3.14}


func main()  {
	x:=3.4
	fmt.Println("type",reflect.TypeOf(x))
	v:=reflect.ValueOf(x)
	fmt.Println(v,v.Type(),v.Kind(),v.Float(),v.Interface())//3.4 float64 float64 3.4 3.4
	//通过反射改变值
	fmt.Println(v.CanSet()) //false
	v=reflect.ValueOf(&x)
	fmt.Println(v.CanSet())//false
	v=v.Elem()
	fmt.Println(v.CanSet())//true
	v.SetFloat(4.18)
	fmt.Println(v.Interface())
	fmt.Println("-----------------")

	//-----------反射结构体----------

	value:=reflect.ValueOf(secret)
	typ:=reflect.TypeOf(secret)

	fmt.Println(typ,value.Kind())
	fmt.Println(value.FieldByName("a"))


	for i:=0;i<value.NumField();i++{
		fmt.Println(value.Field(i))
	}
	fmt.Println(typ.NumMethod()) //1 只会得到可导出分方法 Intro
	fmt.Println(value.Method(0).Call(nil))



	





}

