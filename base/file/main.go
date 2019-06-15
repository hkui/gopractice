package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)


type s struct {
	nums []int
	len int
}
func main() {




	var arr [10]s


	for i:=0;i<10;i++{

		arr[i].nums=append(arr[i].nums,i)
	}
	fmt.Println(arr[1].nums)

	os.Exit(1)



	tm2, _ := time.ParseInLocation("2006-01-02 15:04:05", "2015-02-18 00:00:00",time.Local)
	tm3, _ := time.ParseInLocation("2006-01-02 15:04:05", "2019-02-18 18:01:02",time.Local)

	fmt.Println(tm2.Unix(),"\n",tm3.Unix())



	/*var s  = time.Unix(1560178259, 0)
	println(s.Year(),s.Month(),s.Day(),s.Hour(),s.Minute(),s.Second())*/

	os.Exit(1)

	path,err:=filepath.Abs("./")
	if err!=nil{
		panic(err)
	}
	file:=path+"/base/file/user_consume.txt"

	fileh,err:=os.Open(file)
	if err!=nil{
		panic(err)
	}
	defer  fileh.Close()
	var n  int
	for  {
		i, err := fmt.Fscanf(fileh, "%d\n", &n)
		if i== 0 || err!=nil{
			break
		}
		fmt.Printf("i=%d,num=%d\n",i,n)

	}




}
