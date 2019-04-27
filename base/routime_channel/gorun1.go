package main

import (
	"fmt"
	"time"
)

func longWait()  {
	fmt.Println("begin longWait")
	time.Sleep(5*1e9)// sleep 5 seconds
	fmt.Println("end longWait")
}
func shortWait()  {
	fmt.Println("begin shortWait ")
	time.Sleep(2*1e9)// sleep 5 seconds
	fmt.Println("end shortWait")
}
func main()  {
	fmt.Println("In main")
	go longWait()
	go shortWait()
	fmt.Println("about to sleep in main")
	time.Sleep(10*1e9)
	fmt.Println("end main")


}
