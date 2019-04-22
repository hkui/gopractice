package main

import (
	"fmt"
	"regexp"
)

func main()  {
	str:="abc12b2c123F1g1E12php,python,PhP"

	p:=regexp.MustCompile(`(?i:[a-z]+?[\d+])`)

	r:=p.FindAll([]byte(str),10)
	for _,v:=range r{
		fmt.Println(string(v))
	}


/*
	pattern:="PHP"
	//r 始终是nil,表示是否出现err
	if ok,_:=regexp.Match(pattern,[]byte (str));ok{
		fmt.Println("find")
		fmt.Println(ok)

	}else{
		fmt.Println(ok)
		fmt.Println("not find")
	}
*/






}





