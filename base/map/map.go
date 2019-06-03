package main

import "fmt"
/**
查找不重复字串的最大长度
 */
func diffStrMaxLen(s string) int {
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

func main()  {
	fmt.Println(diffStrMaxLen("abcbdefg"))//6
	fmt.Println(diffStrMaxLen("aaa"))//1
	fmt.Println(diffStrMaxLen("abcdefabc"))//6
	fmt.Println(diffStrMaxLen("你好吗你好啊yes"))//7
}
