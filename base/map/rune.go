package main

import "fmt"

func maxDiffLen()  {

}

func main()  {

	var s string="abcde";
	s1:=[]byte(s)

	fmt.Println(s1)

	for i,v:=range s1 {
		fmt.Println(i,v)
	}
}


