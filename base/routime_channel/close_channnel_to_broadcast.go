package main

import (
	"fmt"
	"math/rand"
	"runtime"
)

func GenerateIntA(done chan struct{}) chan int {
	ch := make(chan int)
	go func() {
	Label:
		for {
			select {
			case ch <- rand.Int():
				//增加1路监听，即为对退出通知信号的监听
			case <-done:
				break Label
			}
		}
	 //收到通知后关闭通道
		close(ch)
	}()
	return ch
}
func main() {
	done:=make(chan  struct{})
	ch:=GenerateIntA(done)

	fmt.Println(<-ch)
	fmt.Println(<-ch)
	//发送通知，告诉生产者停止生产
	close(done)

	fmt.Println(<-ch)
	fmt.Println(<-ch)

	fmt.Println("NumGoroutine=",runtime.NumGoroutine())

}
/**
5577006791947779410
8674665223082153551
0
0
NumGoroutine= 1
 */
