package main

import "fmt"

//chan 参数cap; len()写入长度,cap()获取容量
//map 参数cap; len()写入长度,cap()语法错误
//slice 参数len,cap, len()写入长度,cap()获取容量
func main() {
	//chan type cap
	c := make(chan int, 10)
	fmt.Println(len(c)) //0 查看元素个数,初始化未写入所以为0
	c <- 1
	fmt.Println(len(c)) //1
	fmt.Println(cap(c)) //10

	//map type cap
	m := make(map[int]bool, 1)
	fmt.Println(len(m)) // 0
	m[2] = true
	m[1] = true
	fmt.Println(len(m)) // 2
	//fmt.Println(cap(m)) //语法错误,cap函数的参数m无效

	//slice len cap
	s := make([]int, 0, 10)
	fmt.Println(len(s)) // 0
	s = append(s, 1)
	fmt.Println(len(s)) // 1
	fmt.Println(cap(s)) // 10
}
