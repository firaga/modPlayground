package main

import (
	"fmt"
	"math/rand"
)

/*
注意点:
1. new的seed非线程安全
2. 多协程new时使用相同的种子,会导致大量线程算出一样的分布,在服务发现时对某台机器造成过大负载
3. 不同种子，随机序列发生碰撞的概率高于单个碰撞概率的乘积。 这是因为存在生日问题。(概率计算问题)

https://www.haohongfan.com/docs/gohandbook/rand-chapter/2021-01-23-rand/
https://cloud.tencent.com/developer/article/1911302
*/
func main() {
	rand.Seed(1)
	rand.Intn(1)
	r := rand.New(rand.NewSource(11))
	n := 10
	for i := 0; i < 100; i++ {
		a := r.Intn(n)
		fmt.Println(a)
	}
}
