package main

//命令行参数获取的2种方式
import (
	"flag"
	"fmt"
	"os"
)
var name string
func init(){
	flag.StringVar(&name,"name","","The greeting object.")
}
func main()  {
	if len(os.Args)>1{
		for _,v:=range os.Args{
			fmt.Println(v)
		}

	}
	fmt.Println("----------")
	flag.PrintDefaults()
	flag.Parse()
	if name==""{
		fmt.Println("请指定name")
		return
	}
	fmt.Printf("hello ,%s",name);
}

