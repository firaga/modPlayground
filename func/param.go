package main

import (
	"fmt"
)

func main() {
	cha := make(chan int)
	//playground
	a := 1
	c := 2

	go r2(func() {
		a += 10
		fmt.Printf("func a=%v,ptr=%d\n", a, &a)
		cha <- 1
	})
	<-cha
	a += 100
	fmt.Printf("main a=%v,ptr=%d\n", a, &a)
	fmt.Printf("main c=%v,ptr=%d\n", c, &c)
}

func r2(f func()) {
	fmt.Println("start goroutine 2")
	defer fmt.Println("end goroutine 2")
	go f()
}
