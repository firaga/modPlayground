package main

import (
	"fmt"
	"time"
)

func main() {
	go p()

	fmt.Println("main sleep 10s")
	time.Sleep(10 * time.Second)
	fmt.Println("main end")
}

func p() {
	defer func() {
		fmt.Println("defer p")
	}()
	fmt.Println("run p")
	go s()
	fmt.Println("end p")
}

func s() {
	fmt.Println("run s 2s")
	time.Sleep(2 * time.Second)
	fmt.Println("end s")
}

func gs() {

}
