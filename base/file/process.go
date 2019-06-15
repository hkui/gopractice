package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

var format string="2006-01-02 15:04:05"
var timestrStart="2019-05-11 00:00:00"
var timestrEnd="2019-05-12 00:00:00"
var diff int64 =60  //1分钟
type res struct {
	nums [100]int64
	count int
}
func main() {

	var countArr [2000]res

	timestampStart,_:=time.ParseInLocation(format,timestrStart,time.Local)
	timestampEnd,_:=time.ParseInLocation(format,timestrEnd,time.Local)

	tstart:=timestampStart.Unix()
	tend:=timestampEnd.Unix()
	fmt.Println(tstart,tend)
	arr:=readFromfile()

	for i:=tstart;i<tend;{
		left:=i
		right:=left+diff
		for k,tt:=range arr {
			if tt >= left && tt<right {
				countArr[i].count++

				//countArr[i].nums=append(countArr[i].nums,tt)


				arr=append(arr[:k],arr[k+1:]...)
			}
		}
		i+=diff
	}

	for k,v:=range countArr{
		fmt.Printf("%d=>%d\n",k,v)
	}







}
func readFromfile()[]int64  {
	var arr []int64
	path,_:=filepath.Abs("./")

	file:=path+"/base/file/1.txt"

	fileh,err:=os.Open(file)
	if err!=nil{
		panic(err)
	}
	defer  fileh.Close()
	var n  int64
	for  {
		i, err := fmt.Fscanf(fileh, "%d\n", &n)
		if i== 0 || err!=nil{
			break
		}
		arr= append(arr, n)

	}
	return arr
}
