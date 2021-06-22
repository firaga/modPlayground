package main

import "time"

func foo() {
	i := 1
	go func() {
		i++
	}()
	time.Sleep(time.Second)
	println(i)
}

func bar() {
	i := 1
	go func(i int) {
		i++
	}(i)
	time.Sleep(time.Second)
	println(i)
}
func main() {
	//闭包引用
	//php没有闭包引用?
	foo()
	//变量拷贝
	bar()
}

//作者：尼不要逗了
//链接：https://zhuanlan.zhihu.com/p/63354092
//来源：知乎
//著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
