package main

import (
	"./duckType"
	"fmt"
)

func main()  {
	/*
	b:=new (duckType.Bird)
	duckType.DuckDance(b)
	*/
	r:=duckType.Rectangle{5,3}
	q:=&duckType.Square{5}
	shapes:=[]duckType.Shaper{r,q}

	for n,_:=range shapes{
		fmt.Println("shapes=",shapes[n])
		fmt.Println("area shapes=",shapes[n].Area())
	}




}
