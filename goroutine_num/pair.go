package main

import (
	"fmt"
	"time"
)

func main() {
	go parent()
	time.Sleep(5 * time.Second)
}
func parent() {
	defer fmt.Println("parent end")
	go child()
}
func child() {
	for {
		fmt.Println("child")
		time.Sleep(1 * time.Second)
	}
}
