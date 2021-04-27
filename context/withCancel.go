package main

import (
	"time"
	"runtime"
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()
	cancelCtx, cancelFunc := context.WithCancel(ctx)
	go task(cancelCtx)
	time.Sleep(time.Second * 3)
	cancelFunc()                                                 // 取消 context
	time.Sleep(time.Second * 1)                                  // 延时等待协程退出
	fmt.Println("number of goroutine: ", runtime.NumGoroutine()) // 协程数量
}

func task(ctx context.Context) {
	newCtx, canelFunc := context.WithCancel(ctx)
	go task2(newCtx)
	i := 1
	for {
		select {
		case <-ctx.Done(): // 接收取消信号
			fmt.Println("Gracefully exit")
			fmt.Println(ctx.Err()) // 取消原因
			return
		default:
			fmt.Println(i)
			time.Sleep(time.Second * 1)
			i++
		}
	}
	fmt.Println("forloop end")
	defer canelFunc()
}
func task2(ctx context.Context) {

	i := 1
	for {
		select {
		case <-ctx.Done(): // 接收取消信号
			fmt.Println("Gracefully2 exit")
			fmt.Println(ctx.Err()) // 取消原因
			return
		default:
			fmt.Println(i)
			time.Sleep(time.Second * 1)
			i++
		}
	}
}
