package main
import "fmt"
func push(ch chan int)  {
	for i:=1;;i++{
		fmt.Println("[push start ]",i)
		ch<-i
		fmt.Println("[push end]",i)
	}
}
func pop(ch chan int,done chan int)  {
	i:=0
	for  {
		i++
		if i>4 {
			fmt.Println("break")
			done<-1
			break
		}
		fmt.Println("pop start ",i)
		fmt.Println(<-ch,"pop",i)
		fmt.Println("pop end ",i)
	}
}
func main()  {
	ch:=make(chan int)
	done:=make(chan int)
	go push(ch)
	go pop(ch,done)
	<-done
}
