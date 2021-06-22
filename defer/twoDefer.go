package main

import (
	"fmt"
	"runtime"
)

func main() {
	defer fmt.Println("hello1")
	defer fmt.Println("hello2")

	fmt.Println("world")
	runtime.Breakpoint()
}

//作者：尼不要逗了
//链接：https://zhuanlan.zhihu.com/p/63354092
//来源：知乎
//著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
