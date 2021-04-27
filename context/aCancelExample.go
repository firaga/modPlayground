package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

var c = 1

func doSome(i int) error {
	c++
	fmt.Println(c)
	if c > 3 {
		return errors.New("err occur")
	}
	return nil
}

func speakMemo(ctx context.Context, cancelFunc context.CancelFunc) {
	for {
		select {
		case v := <-ctx.Done():
			fmt.Printf("ctx.Done %v %t", v, v)
			return
		default:
			fmt.Println("exec default func")
			err := doSome(3)
			if err != nil {
				fmt.Printf("cancelFunc()")
				cancelFunc()
			}
		}
	}
}

func main() {
	rootContext := context.Background()
	ctx, cancelFunc := context.WithCancel(rootContext)
	go speakMemo(ctx, cancelFunc)
	time.Sleep(time.Second * 5)
}
