/*
1.数组的初始化和遍历 ,二维数组存储的类型必须一致
2.for  range  遍历数组

*/

package main

import "fmt"
func main(){
	//arr:=[...]int {1,2,3,5,7}
	arr1:=[2][3]int{{2,3,1},{4,5,7}}
	
	/*
	for i:=0;i<len(arr);i++{
		fmt.Printf("%d",arr[i])
	}
	fmt.Printf("--------------\n")
	*/
	
	//取得索引和值
	/*
	for i,v:=range arr{
		fmt.Printf("%d->%d\n",i,v)
	}
	*/

	// fmt.Printf("%d->%d\n",i,v) //继续遍历，i,v为undefined
	
	
	/* 忽略值
	for i:=range arr{
		fmt.Printf("i=%d\n",i)
	}
	*/
	//忽略索引
	/*
	for _,v:=range arr{
		fmt.Printf("v=%d\n",v)
	}
	*/
	
	for i,v:=range arr1{
		for ii,vv:=range v{
			fmt.Printf("i=%d,ii=%d,vv=%d\n",i,ii,vv)
		}
	}
	arr2:=[2][3]string{{"a","b","c"},{"aa","bb","cc"}} 
	for i,v:=range arr2{
		for ii,vv:=range v{
			fmt.Printf("i=%d,ii=%d,vv=%s\n",i,ii,vv)
		}
	}
	
	
	
}