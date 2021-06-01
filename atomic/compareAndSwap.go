package main

import (
	"sync/atomic"
	"unsafe"
)

func main() {
	var p unsafe.Pointer
	i := 1
	p = unsafe.Pointer(&i)
	newP := 42
	ret := atomic.CompareAndSwapPointer(&p, nil, unsafe.Pointer(&newP))

	v := (*int)(p)
	println(*v)
	println(ret)
}
