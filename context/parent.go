package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		fmt.Println("dad start")
		//父协程
		defer func() {
			fmt.Println("exit dad")
		}()
		go func() {
			fmt.Println("kid start")
			//子协程
			defer func() {
				fmt.Println("exit kid")
			}()
			time.Sleep(2 * time.Second)
			fmt.Println("kid end")
		}()
		time.Sleep(1 * time.Second)
		fmt.Println("dad end")
	}()
	time.Sleep(5 * time.Second)
}
