package main

import (
	"fmt"
	"sort"
)

type sorter interface{
	len()int
	swap(i ,j int)
	less(i,j int)bool

}
type IntArray []int
func (p IntArray) Len() int           { return len(p) }
func (p IntArray) Less(i, j int) bool { return p[i] < p[j] }
func (p IntArray) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func main()  {
	data:=[]int{2,1,5,3,6,9,0}
	a := sort.IntArray(data) //conversion to type IntArray from package sort
	sort.Sort(a)



}

