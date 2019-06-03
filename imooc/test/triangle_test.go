package main

import "testing"

func TestTriangle(t *testing.T) {
	tests:=[]struct{a,b,c int}{
		{3,4,5},
		{5,12,13},
		{30,40,50},
	}
	for _,tt:=range tests  {
		if actual:=Triangle(tt.a,tt.b);actual!=tt.c{
			t.Errorf("Triangle(%d,%d)=%d,test get %d\n",tt.a,tt.b,tt.c,actual)
		}
	}
}