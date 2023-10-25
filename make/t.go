package main

import (
	"sync/atomic"
	"time"
)

// Spin Spin是一个锁变量，实现了Lock和Unlock方法
type Spin int32

func (l *Spin) Lock() {
	// 原子交换，0换成1
	for !atomic.CompareAndSwapInt32((*int32)(l), 0, 1) {
	}
}

func (l *Spin) Unlock() {
	// 原子置零
	atomic.StoreInt32((*int32)(l), 0)
}

type Locker interface {
	Lock()
	Unlock()
}

func main() {
	var l Locker
	l = new(Spin)
	var n int
	// 两个协程
	for i := 0; i < 2; i++ {
		go routine(i, &n, l, 200*time.Millisecond)
	}
	select {}
}

func routine(i int, v *int, l Locker, d time.Duration) {
	// 实现自旋加锁
	for {
		func() {
			l.Lock()
			defer l.Unlock()
			*v++
			println(*v, i)
			time.Sleep(d)
		}()
	}
}
