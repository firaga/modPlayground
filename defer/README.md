
[深入理解 Go 语言 defer](https://www.zhihu.com/search?type=content&q=go%20%20defer)
newChangeRet.go

[go中defer实现原理](https://juejin.cn/post/6975686540601245709?utm_source=gold_browser_extension#comment)
changeRet.go
闭包
foobar.go
closure.php

go run -gcflags="-S" twoDefer.go
run.asm

go build -gcflags '-l' -o defer twoDefer.go
go tool objdump -s "main\.main" defer
o.asm


go test -bench=. b_test.go -run=none
对比差距不大
BenchmarkCall
BenchmarkCall-16                100000000               10.79 ns/op
BenchmarkDeferCall
BenchmarkDeferCall-16           92463784                12.56 ns/op

其他
1. 1.16的表现和文章中不一致
   deferproc应该在defer处执行,1.16没有这个方法;
   deferreturn文章中执行了两次,1.16执行了1次.
   同时性能如上,差距不大.
   

q1. deferreturn的作用
q2. 和panic的关系

todo
1. [深入理解defer（下）defer实现机制](https://zhuanlan.zhihu.com/p/69455275)
2. [深入理解defer（上）defer基础](https://zhuanlan.zhihu.com/p/68702577)
3. [Go语言panic/recover的实现](https://zhuanlan.zhihu.com/p/72779197)