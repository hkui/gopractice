package main

import "fmt"

func sum1(a []int, c chan int) {
	total := 0
	for _, v := range a {
		total += v
	}
	c <- total  // send total to c
	fmt.Println(total);
}

func main() {
	a := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)

	go sum1(a[:len(a)/2], c) //7,2,8
	go sum1(a[len(a)/2:], c) //-9,4,0
	x, y := <-c, <-c  // receive from c

	fmt.Println(x, y)
}