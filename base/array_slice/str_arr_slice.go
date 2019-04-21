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

	//两个切片指向同一数组，互相改变了
	/*
	s:=[]byte{'a','b','c','d'}
	s1:=s[:]
	s1[0]='e'
	s2:=s[:]
	s2[1]='f'
	fmt.Println(string(s),string(s1),string(s2)) //均是efcd
	*/

	//删除索引位于1的元素
	/*
	str:=[]byte{'p','a','h','p'}

	str1:=append(str[:1],str[2:]...)

	fmt.Println(string(str1))
	 */

	//删除切片a中i到j的位置的元素

	/*
	a:=[]byte{'a','b','c','d','e','f'}
	a=append(a[:2],a[4:]...)
	print(string(a))//abef
	*/

	//为切片扩展j个元素的长度
	/*
	var a=[]int{1,3,5,7}
	a=append(a,make([]int,10)...)
	fmt.Println("a=%d",a)
	*/

	//在索引i插入元素x
	/*
	a:=[]byte{'a','b','d','e'}
	a=append(a[:2],append([]byte{'c'},a[2:]...)...)
	print(string(a));//abcde
	*/

	//在i的位置插入长度为j的新切片
	/*
	a:=[]byte{'a','b','d','e'}
	a=append(a[:2],append([]byte{'m','n'},a[2:]...)...)
	print(string(a))//abmnde
	 */

}


