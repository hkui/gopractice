package main

import "math"
//勾股定理
func Triangle(a,b int)int  {
	return int(math.Sqrt(float64(a*a+b*b)))
}
//查找不重复字串的最大长度
func DiffStrMaxLen(s string) int {
	lastOccured:=make(map[rune]int)
	len:=0
	start:=0
	rs:= []rune (s);
	for i,ch:=range rs{
		if lastIndex,exists:=lastOccured[ch];exists && lastIndex>=start{
			start=lastIndex+1
		}
		if i-start+1>len{
			len=i-start+1
		}
		lastOccured[ch]=i
	}
	return len
}

func main() {
	
}
