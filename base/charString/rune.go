package main

import (
	"fmt"
	"unicode/utf8"
)

func main()  {
	s:="php最好p"
	bs:=[]byte(s)

	fmt.Println("len=",len(s)) //字节长度
	fmt.Println("RuneCountInString=",utf8.RuneCountInString(s))
	fmt.Printf("%X\n",s)

	for i,v:=range s{  //v is rune
		fmt.Printf("%d=>%X ",i,v)
	}
	fmt.Println("\n---------------------------")

	for i,v:=range bs{
		fmt.Printf("%d=>%X ",i,v)
	}

	ch,size:=utf8.DecodeRune(bs)
	fmt.Println()

	fmt.Println("ch=",ch,"size=",size)

	for i,ch:=range []rune(s){
		fmt.Printf("%d=>%c ",i,ch)
	}

}

