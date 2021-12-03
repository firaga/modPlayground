//作者：qiya
//链接：https: //www.zhihu.com/question/450188866/answer/1792300859
//来源：知乎
//著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

package main

import (
	"sync"
)

func main() {
	// channel 初始化
	c := make(chan int, 10)
	// 用来 recevivers 同步事件的
	wg := sync.WaitGroup{}

	// sender（写端）
	go func() {
		// 入队
		c <- 1
		c <- 2
		// ...
		// 满足某些情况，则 close channel
		close(c)
	}()

	// receivers （读端）
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// ... 处理 channel 里的数据
			for v := range c {
				//fmt.Println(v)
				_ = v
			}
		}()
	}
	// 等待所有的 receivers 完成；
	wg.Wait()
}
