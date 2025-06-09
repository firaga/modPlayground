package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	// 创建一个广播通道
	broadcast := make(chan time.Time)

	// 广播 Ticker 事件
	go func() {
		for t := range ticker.C {
			// 将事件广播到所有消费者
			broadcast <- t
		}
	}()

	// Goroutine 1
	go func() {
		for t := range broadcast {
			fmt.Println("Goroutine 1 received tick at", t)
		}
	}()

	// Goroutine 2
	go func() {
		for t := range broadcast {
			fmt.Println("Goroutine 2 received tick at", t)
		}
	}()

	// 主 Goroutine 等待一段时间以观察输出
	time.Sleep(5 * time.Second)
}
