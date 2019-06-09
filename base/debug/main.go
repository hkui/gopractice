package main

import (
	"fmt"
	"time"
)

func main() {
	for i:=0;i<5;i++{
		j:=i
		go func() {
			fmt.Println(j)
		}()
	}
	time.Sleep(1*time.Second)
}
