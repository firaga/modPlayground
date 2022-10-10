package main

import (
	"sync"
	"sync/atomic"
)

var numCalcsCreated2 int32

// 创建实例的函数
func createBuffer2() interface{} {
	// 这里要注意下，非常重要的一点。这里必须使用原子加，不然有并发问题；
	atomic.AddInt32(&numCalcsCreated2, 1)
	buffer := make([]byte, 1024)
	return &buffer
}

func main() {
	bufferPool := &sync.Pool{
		New: createBuffer2,
	}
	buffer := bufferPool.Get()
	bufferPool.Put(buffer)
	buffer = bufferPool.Get()
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		buffer := bufferPool.Get()
		bufferPool.Put(buffer)
		wg.Done()
	}()
	wg.Wait()

}
