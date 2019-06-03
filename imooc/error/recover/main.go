package main

import (
	"fmt"
)

func tryRecover() {
	defer func() {
		r := recover()
		if r != nil {
			if err, ok := r.(error); ok {
				fmt.Printf("error=%s\n", err)
			}
			fmt.Println("#######all####################")
			fmt.Printf("%T====%v\n", r, r)
		} else {
			fmt.Println("no errors,no bugs!")
		}

	}()
	a := 0
	fmt.Println(10 / a)
	//panic("hahah")
	//panic(errors.New("哈哈"))

}
func main() {
	tryRecover()
}
