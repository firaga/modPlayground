[字节跳动 Go 语言面试高频题 01：内存分配](https://zhuanlan.zhihu.com/p/352133292?utm_source=wechat_session&utm_medium=social&utm_oi=901408282558742528)
#小内存分配
mspan -> mcentral -> heap 申请span,申请过多时返回连续大内存arena
heap应该是管理arena的结构
#大内存分配
mheap分配page(page应该是大于arena的层级结构?)


![](v2-96956b2a183e956855d1551f76f77b71_720w.jpeg)