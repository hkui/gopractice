package main

import "testing"

func TestDiffStrMaxLen(t *testing.T) {
	tests:=[]struct{
		str string
		len int
	}{
		{"aaa",1},
		{"abcdab",4},
		{"加油go加油！",5},

	}
	for _,tt:=range tests{
		if actual:=DiffStrMaxLen(tt.str);actual!=tt.len{
			t.Errorf("DiffStrMaxLen(%s)=%d,excepted %d\n",tt.str,actual,tt.len)
		}
	}
}

func BenchmarkDiffStrMaxLen(b *testing.B) {
	s:="加油go加油！"
	len:=5

	for i:=0;i<b.N ;i++  {
		actual:=DiffStrMaxLen(s)
		if actual!=len {
			b.Errorf("DiffStrMaxLen(%s)=%d,excepted %d\n",s,actual,len)
		}
	}
}