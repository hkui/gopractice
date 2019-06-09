package main
import "fmt"
func push(ch chan int)  {
	for i:=1;;i++{
		fmt.Println("[push start ]",i)
		ch<-i
		fmt.Println("[push end]",i)
	}
}
func pop(ch chan int)  {
	i:=0
	for  {
		i++
		if i>4 {
			fmt.Println("break")
			break
		}
		fmt.Println("pop start ",i)
		fmt.Println(<-ch,"pop",i)
		fmt.Println("pop end ",i)
	}
}
func main()  {
	ch:=make(chan int)
	go push(ch)
	pop(ch)
}/*
pop start  1
[push start ] 1
[push end] 1
[push start ] 2
1 pop 1
pop end  1
pop start  2
2 pop 2
pop end  2
pop start  3
[push end] 2
[push start ] 3
[push end] 3
[push start ] 4
3 pop 3
pop end  3
pop start  4
4 pop 4
pop end  4
break
*/
