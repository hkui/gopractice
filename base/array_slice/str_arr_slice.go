/**
1.改变字符串的值
2.append操作
 */

package main

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
	/*
	str:=[]byte{'p','a','h','p'}

	str1:=append(str[:1],str[2:]...)

	fmt.Println(string(str1))
	 */

	/*
	s:=[]byte{'a','b','c','d'}
	s1:=s[:]
	s1[0]='e'
	s2:=s[:]
	s2[1]='f'
	fmt.Println(string(s),string(s1),string(s2)) //均是efcd
	*/











}

