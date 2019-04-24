package main

import "fmt"

type Shaper interface {
	Area() float32
}
type Square struct {
	side float32
}

func (sq * Square)Area() float32{
	return sq.side*sq.side
}

type Rectangle struct {
	length,width float32
}

func (r Rectangle)Area()float32  {
	return r.length*r.width
}

func main()  {
	sq1:=new(Square)
	sq1.side=10

	fmt.Println(sq1.Area())

	var areainterf Shaper
	areainterf=sq1
	fmt.Println(areainterf.Area())

	sq2:=&Rectangle{3,8}
	fmt.Println(sq2.Area())

	shaps:=[]Shaper{sq1,sq2}
	for k,_:=range shaps{
		fmt.Println(shaps[k],shaps[k].Area())
	}


}

