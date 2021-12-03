//作者：qiya
//链接：https: //www.zhihu.com/question/450188866/answer/1792300859
//来源：知乎
//著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

package main

import (
	"context"
	"sync"
	"time"
)

func main() {
	// channel 初始化
	c := make(chan int, 10)
	// 用来 recevivers 同步事件的
	wg := sync.WaitGroup{}
	// 上下文
	ctx, cancel := context.WithCancel(context.TODO())

	// 专门关闭的协程
	go func() {
		time.Sleep(2 * time.Second)
		cancel()
		// ... 某种条件下，关闭 channel
		close(c)
	}()

	// senders（写端）
	for i := 0; i < 10; i++ {
		go func(ctx context.Context, id int) {
			select {
			case <-ctx.Done():
				return
			case c <- id: // 入队
				// ...
			}
		}(ctx, i)
	}

	// receivers（读端）
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// ... 处理 channel 里的数据
			for v := range c {
				_ = v
			}
		}()
	}
	// 等待所有的 receivers 完成；
	wg.Wait()
}
