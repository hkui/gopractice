package main

import (
	"bufio"
	"fmt"
	"os"
)

func tryDefer()  {
	for i:=0;i<10 ;i++  {
		defer fmt.Printf("defer i=%d\n",i)
		fmt.Printf("i=%d\n",i)
		if i==5{
			panic("print too many")
		}
	}
}
func writeFile(filename string) {
	file,err:=os.Create(filename)
	if err!=nil{
		panic(err)
	}
	defer file.Close()
	writer:=bufio.NewWriter(file)
	defer writer.Flush()
	for i:=0;i<10 ;i++  {
		fmt.Fprintln(writer,i)
	}
}
func writeFile1(filename string)  {
	file,err:= os.OpenFile(filename, os.O_EXCL|os.O_CREATE|os.O_TRUNC, 0666)
	if err!=nil{
		if pathError,ok:=err.(*os.PathError);!ok{
			panic(err)
		}else {
			fmt.Printf("%s,%s,%s\n",pathError.Op,pathError.Path,pathError.Err)//open,aa.txt,The file exists
		}
	}
	defer file.Close()
	writer:=bufio.NewWriter(file)
	defer writer.Flush()
	for i:=0;i<10 ;i++  {
		fmt.Fprintln(writer,i)
	}


}

func main()  {
	//writeFile("aa.txt")
	writeFile1("aa.txt")
}
