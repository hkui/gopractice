package main

import "fmt"

type Human struct {
	name string
	age int
	phone string
}

type Student struct {
	Human //匿名字段
	school string
}

type Employee struct {
	Human //匿名字段
	company string
}

//在human上面定义了一个method
func (h *Human) SayHi() {
	fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}
func (e *Employee) SayHi() {
	fmt.Printf("Hi, I am %s ,I work at %s ,you can contact me on %s\n",e.name,e.company,e.phone)
	e.Human.SayHi()
}

func main() {
	mark := Student{Human{"Mark", 25, "74110"}, "MIT"}
	sam := Employee{Human{"Sam", 45, "111-888-999"}, "Golang Inc"}

	mark.SayHi()
	sam.SayHi()
}