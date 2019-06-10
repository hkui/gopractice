package main

import (
	"fmt"
	"time"
)

func main() {
	tick1 := time.Tick(time.Second)
	tick2 := time.Tick(2 * time.Second)
	after := time.After(10 * time.Second)
	t2i := 1
	//EXIT:
	for {
		select {
		case t1 := <-tick1:
			fmt.Println("tick1", t1)
		case t2 := <-tick2:
			if t2i > 4 {
				break //只是跳出了select,以后的select不在监控这个分支
			}
			fmt.Println("tick2", t2, t2i)
			t2i++
		case a1 := <-after:
			fmt.Println("after ", a1)
			//break EXIT
			goto END

		}
	}
END:
}
