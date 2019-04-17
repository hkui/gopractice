/*
1.从数组创建切片,切片与数组
2.切片传递给函数
3.无相关数组时用make,new 创建切片
4.切片重新分片
5.切片的复制与追加
*/
package main

import "fmt"

func sum(a []int ) int{
	total:=0
	for _,v:=range a{
		total+=v
	}
	return total

}
//https://github.com/Unknwon/the-way-to-go_ZH_CN/blob/master/eBook/07.5.md
func AppendByte(slice []byte,data ... byte)[]byte{
	sliceLen:=len(slice)
	sliceCap:=cap(slice)
	dataLen:=len(data)
	if sliceCap<sliceLen+dataLen {
		
	}
	copy(slice[len(slice):],data)
	
	return slice
}

func main(){
	/*
	arr:=[]int{1,3,5,7,9,2};
	s:=arr[1:3] //[1,3)前闭后开
	for _,v:=range s{
		fmt.Printf("v=%d\n",v)
	}
	fmt.Printf("len(s)=%d,cap(s)=%d\n",len(s),cap(s)) //2,5
	fmt.Printf("sum=%d",sum(s))
	*/
	
	/*
	s1:=make([]int ,5) //初始长度为5 len=5,cap=5
	
	s2:=make([]int,5,10) //len(s2)=5,cap(s2)=10
	s3:=new([10]int)[0:5]
	
	fmt.Printf("len(s1)=%d,cap(s1)=%d\n",len(s1),cap(s1))
	fmt.Printf("len(s2)=%d,cap(s2)=%d\n",len(s2),cap(s2))
	fmt.Printf("len(s3)=%d,cap(s3)=%d\n",len(s3),cap(s3))
	*/
	/*
	//切片重新分片
	arr:=[10]int{0,1,2,3,4,5,6,7}
	slice1:=arr[5:7] //5,6
	slice2:=slice1[:4] //5,6,7,0
	fmt.Printf("len(slice1)=%d,cap(slice1)=%d\n",len(slice1),cap(slice1)) //2,5
	fmt.Printf("len(slice2)=%d,cap(slice2)=%d\n",len(slice2),cap(slice2)) //4,5
	*/
	//copy与append
	
	sl_from:= []int{1,3,5}
	sl_to:=make([]int,10)
	sl_to[0]=9
	
	n:=copy(sl_to[1:],sl_from)
	fmt.Printf("n=%d,sl_to=%d\n",n,sl_to)//n=3,sl_to=[9 1 3 5 0 0 0 0 0 0]
	
	sl3:=[]int{4,6,8}
	sl3=append(sl3,1,3,5)
	fmt.Printf("sl3=%d\n",sl3)//sl3=[4 6 8 1 3 5]
	
	
	
	
	
	
	
	
	
	
	

	
	
	
	
	



}


