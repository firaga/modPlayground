package main

import "math"

func main() {
	a, b, c := 2.0, 1.0, 0.0
	x, y := a/c, b/c     // infinity
	n := math.NaN()      // not a number
	m := math.Sqrt(-1.0) // not a number
	println(x == y, m == n)
}

//作者：煎鱼eddycjy
//链接：https: //juejin.cn/post/7159417590270394382
//来源：稀土掘金
//著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
