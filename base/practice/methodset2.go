package main

import "fmt"

type List []int

func (l List)Len()int{
	return len(l)
}
func (l *List)Append(val int)  {
	*l=append(*l,val)
}
type Lener interface {
	Len() int
}

type Appender interface {
	Append(int)
}

func LongEnouth(l Lener) bool  {
	return l.Len()*10>42
}


func CountInto(a Appender,start,end int)  {
	for i:=start;i<=end;i++{
		a.Append(i)
	}
}



func main()  {
	var lst List
	if LongEnouth(lst){
		fmt.Println(lst,"lst is long enough\n")

	}
	plst:=new (List)
	CountInto(plst,1,10)
	if LongEnouth(plst){
		fmt.Println(plst)
	}
}
