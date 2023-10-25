package main

import (
	"context"
	"fmt"
	"time"
)

// 测试cancel之后再监听的结果, 正常
func main() {
	ctx := context.Background()
	parent(ctx)
	time.Sleep(3 * time.Second)
}
func parent(ctx context.Context) {
	defer fmt.Println("parent end")
	fmt.Println("parent start")
	ctx, cancel := context.WithCancel(ctx)
	go son(ctx)
	//time.Sleep(2 * time.Second)
	fmt.Println("parent start cancel")
	cancel()
	fmt.Println("parent end cancel")
}
func son(ctx context.Context) {
	defer fmt.Println("son end")
	fmt.Println("son start")
	//time.Sleep(2 * time.Second)
	fmt.Println("son start forloop")
	for {
		select {
		case a := <-ctx.Done():
			fmt.Println("son receive done!", a)
			return
		default:
			fmt.Println("son default")
			time.Sleep(1 * time.Second)
		}
	}
}
