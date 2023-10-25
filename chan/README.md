[Go语言中如何检测一个channel已经被关闭了？](https://www.zhihu.com/question/450188866/answer/1792300859)

[golang chan 最详细原理剖析，全面源码分析！看完不可能不懂的！](https://zhuanlan.zhihu.com/p/299592156)

## 总结
读已关闭的chan返回类型0值,如果有第二个返回值,返回是否关闭bool值

写已关闭的chan抛出异x,可以defer recover处理,不推荐这种实现,可以通过只有单协程关闭解决问题,其他写协程通过ctx(ctx.Done())+select停止写入