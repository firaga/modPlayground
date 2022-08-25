package main

import (
	"context"
	"fmt"
	"time"
)

// 测试多层cancel是否生效
func main() {
	ctx := context.Background()
	level1(ctx)
}

func level1(ctx context.Context) {
	fmt.Println("level 1 start")
	defer fmt.Println("level 1 exit")
	ctx, cancelFunc := context.WithCancel(ctx)
	go level2(ctx)
	time.Sleep(5 * time.Second)
	cancelFunc()
	time.Sleep(1 * time.Second)
}

func level2(ctx context.Context) {
	fmt.Println("level 2 start")
	defer fmt.Println("level 2 exit")
	for {
		select {
		case <-ctx.Done():
			fmt.Println("level 2 receive done!")
			return
		default:
			//fmt.Println("level 2 default")
			//ctx, _ := context.WithCancel(ctx)
			go level3(ctx)
			time.Sleep(1 * time.Second)
		}
	}
}

func level3(ctx context.Context) {
	fmt.Println("level 3 start")
	defer fmt.Println("level 3 exit")
	for {
		select {
		case <-ctx.Done():
			fmt.Println("level 3 receive done!")
			return
		default:
			//fmt.Println("level 3 default")
			//go level3(ctx)
			time.Sleep(1 * time.Second)
		}
	}
}
