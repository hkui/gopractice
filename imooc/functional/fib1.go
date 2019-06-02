package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

type intGen func()int

func (g intGen) Read(p []byte) (n int, err error) {
	next:=g()
	if next >10000{
		return 0,io.EOF
	}
	s:=fmt.Sprintf("%d\n",next)
	//todo 如果p太小就会有问题
	return strings.NewReader(s).Read(p)
}

func printFileCounts(reader io.Reader)  {
	scanner:=bufio.NewScanner(reader)

	for scanner.Scan(){
		fmt.Println(scanner.Text())
	}

}

func fib1() intGen {
	a,b:=0,1
	return func()int {
		a,b=b,a+b
		return a
	}

}
func main()  {
	f:=fib1()
	printFileCounts(f)

}