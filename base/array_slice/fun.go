package main

import "fmt"


func main()  {
	arr:=[]int{2,1,3,7,8,5,-1,20,9,0}
	arr=mapFunc(f10,arr)
	fmt.Println(arr)

}
/*
str:="abcdefg"
s:=str[len(str)/2:] + str[:len(str)/2] //defgabc
fmt.Println(str[0])
*/

//函数练习

//1.编写一个函数，要求其接受两个参数，原始字符串 str 和分割索引 i，然后返回两个分割后的字符串

func split( str string, i int )(string,string)  {
	return str[:i],str[i:]
}

//2.编写一个函数反转字符串

func reverse(str string)string{
	rs:="";
	var i int =len(str)-1;
	for ; i>=0;i--  {
		rs=rs+string(str[i])
	}
	return string(rs)
}

//3.编写一个程序，要求能够遍历一个数组的字符，并将当前字符和前一个字符不相同的字符拷贝至另一个数组
func findDiffstr(strarr []byte) ([]byte) {
	var before byte=0;
	var newarr []byte;

	for _,v:=range strarr{
		if(v!=before && before!=0){
			newarr=append(newarr,v)
		}
		before=v;
	}
	return newarr
}
//4.编写一个程序，使用冒泡排序的方法排序一个包含整数的切片
func sort_budding(arr []int)([]int)  {

	len:=len(arr)
	for i:=0;i<len-1 ;i++  {
		for j:=i+1;j<len ; j++ {
			if(arr[j]<arr[i]){
				tmp:=arr[i]
				arr[i]=arr[j]
				arr[j]=tmp
			}
		}
	}
	return arr
}
/*
5 编写一个函数 mapFunc 要求接受以下 2 个参数：

一个将整数乘以 10 的函数
一个整数列表
最后返回保存运行结果的整数列表
 */
func f10(num int)int  {
	return num*10
}
func mapFunc(f func(int)int,arr []int )([]int)  {
	for i,v:=range arr{
		arr[i]=f(v)
	}
	return arr
}

