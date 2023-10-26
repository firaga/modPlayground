package main

import (
	"fmt"
	"time"
)

func main() {
	// 延时器1
	go time.AfterFunc(3*time.Second, func() {
		fmt.Println("3s")
	})
	// 延时器2
	go func() {
		for {
			select {
			case <-time.After(5 * time.Second):
				fmt.Println("5s")
				return
			}
		}
	}()
	time.Sleep(100 * time.Second)
	fmt.Println("程序结束")

}

func tick() {
	// 定时器
	go func() {
		//tick := time.Tick(1000 * time.Millisecond)
		//myTimer := time.NewTimer(time.Second * 1)
		for {
			select {
			case a := <-time.Tick(5 * time.Second):
				fmt.Println("5s")
				println(a)
			}
		}
	}()
	time.Sleep(100 * time.Second)
	fmt.Println("程序结束")
}
