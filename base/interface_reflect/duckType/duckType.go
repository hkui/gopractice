package duckType

import "fmt"

type IDuck interface {
	Quack()
	Walk()
}

func DuckDance(duck IDuck)  {
	for i:=0;i<=3;i++{
		duck.Quack()
		duck.Walk()
	}
}
type Bird struct {
	
}

func (b *Bird)Quack()  {
		fmt.Println("Quacking")
}

func (b * Bird)Walk()  {
	fmt.Println("Walking")
}

