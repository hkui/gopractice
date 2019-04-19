/**
1.改变字符串的值
2.append操作
 */

package main

import "fmt"

func main()  {
	//修改字符串

	/*
	s:="hello"
	c:=[]byte(s)
	c[1]='p'
	s=string(c)

	fmt.Println(s);
	*/
	//删除索引位于1的元素
	str:=[]byte{'p','a','h','p'}
	print(cap(str[:1]),"------------",cap(str[2:]),"\n")// 4   2

	str=append(str[:1],str[2:]...)
	print(cap(str),"\n")  //4

	fmt.Println(string(str))



}

