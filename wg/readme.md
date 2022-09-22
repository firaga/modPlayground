如何解决wg超时问题

为什么结构体WaitGroup中的counter,waiter,sema顺序不固定
[Golang WaitGroup 原理深度剖析](https://zhuanlan.zhihu.com/p/344973865)

从图中看出，当 `state1` 是 32 位对齐和 64 位对齐的情况下，`state1` 中每个元素的顺序和含义也不一样:

- *当 `state1` 是 32 位对齐：`state1` 数组的第一位是 sema，第二位是 counter，第三位是 waiter。*
- 当 `state1` 是 64 位对齐：`state1` 数组的第一位是 counter，第二位是 waiter，第三位是 sema。

为什么会有这种奇怪的设定呢？这里涉及两个前提:

前提 1：在 WaitGroup 的真实逻辑中， counter 和 waiter 被合在了一起，当成一个 64 位的整数对外使用。当需要变化 counter 和 waiter 的值的时候，也是通过 atomic 来原子操作这个 64 位整数。但至于为什么合在一起，我们会在下文详细讨论。

前提 2：在 32 位系统下，如果使用 atomic 对 64 位变量进行原子操作，调用者需要自行保证变量的 64 位对齐，否则将会出现异常。golang 的官方文档 [sync/atomic/#pkg-note-BUG](https://link.zhihu.com/?target=https%3A//golang.org/pkg/sync/atomic/%23pkg-note-BUG) 原文是这么说的：

> On ARM, x86-32, and 32-bit MIPS, it is the caller’s responsibility to arrange for 64-bit alignment of 64-bit words accessed atomically. The first word in a variable or in an allocated struct, array, or slice can be relied upon to be 64-bit aligned.

因此，在前提 1 的情况下，WaitGroup 需要对 64 位进行原子操作。根据前提 2，WaitGroup 需要自行保证 `count+waiter` 的 64 位对齐。这也是 WaitGroup 采用 `[3]uint32` 存储变量的目的：