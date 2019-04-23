package main



/*
1.使用包的结构体
2.使用带标签的结构体

*/
import (
	"./stru"
	"fmt"
	"reflect"
)


type tagType struct {
	no int "这是编号"
	name string "这是姓名"
}

func refTag(tt tagType,ix int)  {
	ttType:=reflect.TypeOf(tt)
	ixField:=ttType.Field(ix)
	fmt.Println(ixField.Tag)
	/**
	 这是编号
	 这是姓名
	 */
}
func main()  {
	stru1:=new (stru.Stu)
	stru1.Name="hk"
	stru1.Age=18

	//fmt.Println(stru1) //&{18 hk}

	tt:=tagType{1,"john"}
	for i:=0;i<2;i++{
		refTag(tt,i)
	}



}
